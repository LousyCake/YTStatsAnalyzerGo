# YTStatsAnalyzerGo

YTStatsAnalyzerGo is an open-source YouTube Channel Statistics Analyzer written in Go. It leverages the power of the YouTube Data API v3 to fetch channel statistics and analyze top videos based on views, likes, and comments.

## Table of Contents
- [Introduction](#introduction)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
  - [Fetching Channel Statistics](#fetching-channel-statistics)
  - [Analyzing Top Videos](#analyzing-top-videos)
  - [Saving to CSV](#saving-to-csv)
- [Customization](#customization)
  - [Modifying Output](#modifying-output)
  - [Example: Changing CSV Format](#example-changing-csv-format)
- [Contributing](#contributing)
- [License](#license)

## Introduction

YTStatsAnalyzerGo is designed to provide insights into YouTube channel performance. By utilizing the YouTube Data API v3, it enables users to retrieve channel statistics and analyze top videos for a more in-depth understanding of content engagement.

## Getting Started

### Prerequisites

Before using YTStatsAnalyzerGo, ensure you have:

- Go installed on your machine. Download it from [Here](https://go.dev/doc/install).
- A Google Cloud API key for YouTube Data API v3.

### Installation

Clone the repository:

```
git clone https://github.com/yourusername/YTStatsAnalyzerGo.git
cd YTStatsAnalyzerGo
```
Add your __'API key'__ to the apiKey constant in __'main.go'__:

```
const apiKey = "YOUR_API_KEY"
```

Run the application:

```
go run main.go
```

## Usage
### Fetching Channel Statistics
Run the following command to fetch and display channel statistics for a specified channel:

```
go run main.go
```

### Analyzing Top Videos
YTStatsAnalyzerGo can analyze top videos based on views, likes, and comments. Results can be printed to the console or saved to a CSV file. Examples:

+ Analyzing most viewed videos:

```
go run main.go -top=views -count=10
```

+ Analyzing least viewed videos:

```
go run main.go -top=views -count=10 -asc
```

### Saving to CSV
To save the results to a CSV file, use the __'-csv'__ flag:

```
go run main.go -top=likes -count=5 -csv=output.csv
```

## Customization

### Modifying Output

The output format can be customized by modifying the __'getTopVideosCSV'__ function in __'main.go'__.

### Example: Changing CSV Format

To modify the CSV format, update the following lines in __'getTopVideosCSV'__:

```
result.WriteString("Rank,Title,Views,Likes,Comments\n")
result.WriteString(fmt.Sprintf("%d,\"%s\",%d,%d,%d\n",
    i+1, video.Snippet.Title, video.Statistics.ViewCount, video.Statistics.LikeCount, video.Statistics.CommentCount))
```
## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

YTStatsAnalyzerGo is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


