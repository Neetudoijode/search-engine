package parser

import (
    "github.com/xitongsys/parquet-go/reader"
    "github.com/xitongsys/parquet-go-source/local"
)

type Record struct {
    Message        string
    MessageRaw     string
    StructuredData string
    Tag            string
    Sender         string
    Groupings      string
    Event          string
    EventId        string
    NanoTimeStamp  int64
    Namespace      string
}

func ParseParquetFile(filePath string) ([]Record, error) {
    fr, err := local.NewLocalFileReader(filePath)
    if err != nil {
        return nil, err
    }
    defer fr.Close()

    pr, err := reader.NewParquetReader(fr, new(Record), 4)
    if err != nil {
        return nil, err
    }
    defer pr.ReadStop()

    num := int(pr.GetNumRows())
    records := make([]Record, num)
    if err = pr.Read(&records); err != nil {
        return nil, err
    }

    return records, nil
}
