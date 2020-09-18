package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type budgetCategory string

const (
	autoFuel   budgetCategory = "fuel"
	food       budgetCategory = "food"
	mortgage   budgetCategory = "mortgage"
	repairs    budgetCategory = "repairs"
	insurance  budgetCategory = "insurance"
	utilities  budgetCategory = "utilities"
	retirement budgetCategory = "retirement"
)

var (
	ErrBudgetCategoryNotFound = errors.New(" a budget category is not found ")
)

type Transaction struct {
	id       int
	payee    string
	spent    float64
	category budgetCategory
}

func convertToBudgeCategory(category string) (budgetCategory, error) {
	switch category {
	case "fuel", "gas":
		return autoFuel, nil
	case "food":
		return food, nil
	case "mortgage":
		return mortgage, nil
	case "repairs":
		return repairs, nil
	case "car insurance", "life insurance":
		return insurance, nil
	case "utilities":
		return utilities, nil
	default:
		return "", ErrBudgetCategoryNotFound
	}
}

func writeErrorToLog(msg string, err_arg error, data string, logfile string) error {
	logFile, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer logFile.Close()
	now := time.Now()
	day := strconv.Itoa(int(now.Day()))
	month := strconv.Itoa(int(now.Month()))
	year := strconv.Itoa(int(now.Year()))
	hour := strconv.Itoa(int(now.Hour()))
	minute := strconv.Itoa(int(now.Minute()))
	seconds := strconv.Itoa(int(now.Second()))
	fmt.Println()
	_, err = logFile.WriteString("[" + hour + ":" + minute + ":" + seconds + " " + year + "/" + month + "/" + day + "] " + msg + " " + err_arg.Error() + " " + data + "\n")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func parseBankFile(bankTransactions io.Reader, logFile string) []Transaction {
	reader := csv.NewReader(bankTransactions)
	trxs := []Transaction{}
	header := true
	for {
		trx := Transaction{}
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if !header {
			for idx, value := range rec {
				switch idx {
				case 0:
					trx.id, err = strconv.Atoi(value)
					if err != nil {
						log.Fatal(err)
					}
				case 1:
					trx.payee = strings.Trim(value, " ")
				case 2:
					trx.spent, err = strconv.ParseFloat(strings.Trim(value, " "), 64)
					if err != nil {
						log.Fatal(err)
					}
				case 3:
					trx.category, err = convertToBudgeCategory(strings.Trim(value, " "))
					if err == ErrBudgetCategoryNotFound {
						writeErrorToLog("error converting category - ", err, value, logFile)
					}
				}
			}
			trxs = append(trxs, trx)
		}
		header = false
	}

	return trxs
}

func main() {
	logFilePath := flag.String("l", "", "provide log file with flag -l\n")
	transactionFilePath := flag.String("c", "", "transaction file")
	flag.Parse()
	if *logFilePath == "" || *transactionFilePath == "" {
		fmt.Println("log file and transaction file are required")
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Printf("log file - %s\n", *logFilePath)
	fmt.Printf("transaction file - %s\n", *transactionFilePath)

	transactionFileStat, err := os.Stat(*transactionFilePath)
	if err != nil {
		if os.IsNotExist((err)) {
			fmt.Println(transactionFilePath, ": File does not exist!")
			fmt.Println(transactionFileStat)
		}
		os.Exit(1)
	}

	transactionFile, err := os.Open(*transactionFilePath)
	if err != nil {
		log.Fatal(err)
	}

	trxs := parseBankFile(transactionFile, *logFilePath)
	for indx, value := range trxs {
		fmt.Println("Transaction №", indx)
		fmt.Printf("\t%s\n\t%s\n\t%s\n\n", value.payee, strconv.FormatFloat(value.spent, 'f', 2, 64), string(value.category))
	}
}
