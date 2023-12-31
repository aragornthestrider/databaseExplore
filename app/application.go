package app

import (
	"context"
	"databaseExplore/app/config"
	"databaseExplore/app/database"
	"databaseExplore/app/server"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/google/uuid"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Application struct {
	db          *database.Database
	insertCount int
	Logger      *zap.Logger
}

func (app *Application) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	logger, err := NewLogger("info")
	if err != nil {
		log.Fatal("Error in setting up logger")
		return
	}

	go func() {
		exitSignal := <-exit
		logger.Info("Service shutting down main due to: " + exitSignal.String())
		cancel()
	}()

	configPath := flag.String("config", "/config/config.yaml", "Path of config file")
	config := config.ParseConfig(*configPath, logger)
	log.Println(config)

	server := server.NewServer(ctx, 8080, 10, 10, logger)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		err = server.Run()
		if err != nil {
			logger.Error("Error in running server", zap.Error(err))
		}
	}()

	server.SetHealthy(true)
	server.SetReady(true)

	db := new(database.Database)
	err = db.Connect(fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		config.DatabaseConfig.URL,
		config.DatabaseConfig.Port,
		config.DatabaseConfig.Username,
		config.DatabaseConfig.Password,
		config.DatabaseConfig.Database,
	))
	if err != nil {
		logger.Error("Unable to connect to database", zap.Error(err))
		app.Close()
		return
	}
	logger.Info("Database connected")

	app.db = db
	app.Logger = logger
	app.insertCount = config.InsertCount

	err = app.SetupStorage()
	if err != nil {
		app.Logger.Error("Error in setting up storage", zap.Error(err))
		app.Close()
		return
	}

	err = app.TestStorage()
	if err != nil {
		app.Logger.Error("Error in testing storage", zap.Error(err))
		app.Close()
		return
	}

	wg.Wait()

	app.Close()
}

func (app *Application) SetupStorage() error {
	_, err := app.db.Conn.Exec(`
	CREATE TABLE IF NOT EXISTS TESTTABLE(
		UUIDvalue text PRIMARY KEY
	);`)
	if err != nil {
		app.Logger.Error("Error in creating table", zap.Error(err))
		return err
	}
	app.Logger.Info("Storage setup correctly")
	return nil
}

func (app *Application) TestStorage() error {
	insertStatement := "INSERT INTO TESTTABLE (UUIDvalue) VALUES ($1)"
	startTime := time.Now()
	for i := 0; i < app.insertCount; i++ {
		conn, err := app.db.Conn.Conn(context.TODO())
		if err != nil {
			app.Logger.Error("Error in getting connection", zap.Error(err))
			return err
		}
		_, err = conn.ExecContext(context.TODO(), insertStatement, uuid.New().String())
		if err != nil {
			app.Logger.Error("Error in inserting data", zap.Error(err))
			return err
		}
		conn.Close()
	}
	timeDiff := time.Now().Sub(startTime)
	app.Logger.Info("Summary", zap.String("Record count", fmt.Sprintf("%d", app.insertCount)), zap.String("Time taken", timeDiff.String()))

	return nil
}

func (app *Application) Close() {
	err := app.db.Conn.Close()
	if err != nil {
		app.Logger.Error("Error in closing database connection", zap.Error(err))
	}
	app.Logger.Info("Database connection closed!")

	app.Logger.Sync()
	return
}

func NewLogger(configLogLevel string) (*zap.Logger, error) {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	logLevel := parseLogLevel(configLogLevel)
	core := ecszap.NewCore(encoderConfig, os.Stdout, logLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger, nil
}

func parseLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}
