# GoCrawler

GoCrawler is a web crawling framework written in Go, inspired by Scrapy. It uses the Colly library to handle web scraping and aims to provide a similar structure and command-line interface as Scrapy.

## Getting Started

### Prerequisites

- Go 1.20 or later
- Git

### Installation

#### Clone the repository:

```bash
git clone https://github.com/yourusername/your-go-crawler.git
cd your-go-crawler
```

#### Install the dependencies:

```bash
git clone https://github.com/yourusername/your-go-crawler.git
go mod tidy
```

#### Install the gocrawl command:

```
go install ./cmd/gocrawl
```
### Usage

#### Create a New Project, 
Use the gocrawl startproject command to create a new crawling project:

```bash
gocrawl startproject myproject
```

#### Run the Spider
Navigate to your project directory and use the gocrawl crawl command to run the spider:
```bash
cd myproject
gocrawl crawl
```

## Contributing
We welcome contributions! Please follow these steps to contribute:

- Fork the repository.
- Create a new branch (git checkout -b feature/your-feature).
- Make your changes and commit them (git commit -m 'Add some feature').
- Push to the branch (git push origin feature/your-feature).
- Open a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgements
Colly - Elegant Scraping Framework for Gophers

Cobra - A Commander for modern Go CLI interactions

Inspired by Scrapy - An open source and collaborative web crawling framework for Python.

Reference project layout: https://github.com/golang-standards/project-layout
