package main

import (
	"bytes"
	"fmt"
	"image/png"
	"log"

	"github.com/gen2brain/go-fitz"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	// Abrir o arquivo PDF
	doc, err := fitz.New("teste.pdf")
	if err != nil {
		log.Fatalf("Falha ao abrir o arquivo PDF: %v", err)
	}
	defer doc.Close()

	// Criar um novo cliente tesseract
	client := gosseract.NewClient()
	defer client.Close()

	// Iterar sobre as páginas do PDF
	for n := 0; n < doc.NumPage(); n++ {
		page, err := doc.Image(n)
		if err != nil {
			log.Fatalf("Falha ao obter imagem da página: %v", err)
		}

		var buf bytes.Buffer
		err = png.Encode(&buf, page)
		if err != nil {
			log.Fatalf("Falha ao codificar imagem: %v", err)
		}

		client.SetImageFromBytes(buf.Bytes())

		// Realizar OCR na imagem
		text, err := client.Text()
		if err != nil {
			log.Fatalf("Falha ao realizar OCR na imagem: %v", err)
		}

		// Exibir o texto extraído
		fmt.Printf("Texto da página %d:\n%s\n", n, text)
	}
}
