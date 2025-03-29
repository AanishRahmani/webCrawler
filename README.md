# Web Crawler

## Overview
This is a **concurrent web crawler** implemented in Go. It recursively crawls internal pages of a given website, extracts links, and continues crawling while enforcing domain restrictions and concurrency control.

## Features
- **Recursive Crawling**: Automatically follows internal links and extracts more pages.
- **Domain Restriction**: Ensures that only pages from the same domain are crawled.
- **Concurrency Control**: Limits the number of concurrent requests to prevent overwhelming the server.
- **Duplicate Prevention**: Avoids crawling the same page multiple times.

## How It Works
1. The crawler starts at a given **base URL**.
2. It fetches the HTML content of the page and extracts all internal links.
3. It recursively crawls each extracted link, maintaining concurrency using **goroutines and a WaitGroup**.
4. It avoids revisiting previously crawled pages and ensures that only pages within the same domain are followed.
5. A concurrency limit is enforced using **a buffered channel**, preventing excessive requests.

## Installation
To use this crawler, you need to have Go installed on your system.

1. Clone this repository:
   ```sh
   git clone https://github.com/yourusername/web-crawler.git
   cd web-crawler
   ```
2. Install dependencies (if any, but for now, standard Go packages are used).
3. Build and run the program:
   ```sh
   go run main.go https://example.com
   ```

## Usage
- To start crawling, run the script with a **starting URL**.
- The crawler will print visited pages and follow internal links.
- Modify `concurrencyControl` to adjust the number of parallel requests.

## Future Enhancements
- **Max Pages Limit**: Add a restriction to limit the total number of pages crawled to prevent infinite crawling.
- **Improved URL Normalization**: Enhance handling of relative URLs and duplicate links.
- **Customizable Concurrency**: Allow users to specify concurrency limits via command-line arguments.
- **Output Options**: Save crawled URLs to a file or database for further analysis.

## Contribution
Feel free to fork and improve this project! Suggestions and pull requests are welcome.

## License
This project is licensed under the MIT License.

