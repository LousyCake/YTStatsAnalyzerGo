package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

const apiKey = "**********"

func main() {
	ctx := context.Background()

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}

	part := []string{"id", "snippet", "contentDetails", "statistics"}
	channelStats, err := service.Channels.List(part).Id("**********").Do()
	if err != nil {
		log.Fatalf("Error fetching channel statistics: %v", err)
	}

	if len(channelStats.Items) > 0 {
		channel := channelStats.Items[0]
		fmt.Printf("Channel Title: %s\n", channel.Snippet.Title)

		videos, err := getChannelVideos(service, "UCpDJl2EmP7Oh90Vylx0dZtA")
		if err != nil {
			log.Fatalf("Error fetching channel videos: %v", err)
		}

		var output string

		output += getTopVideosCSV("Most Viewed", videos, 10, sortByViewsDesc)
		output += getTopVideosCSV("Least Viewed", videos, 10, sortByViewsAsc)
		output += getTopVideosCSV("Most Liked", videos, 10, sortByLikesDesc)
		output += getTopVideosCSV("Least Liked", videos, 10, sortByLikesAsc)
		output += getTopVideosCSV("Most Commented", videos, 10, sortByCommentsDesc)
		output += getTopVideosCSV("Least Commented", videos, 10, sortByCommentsAsc)

		fmt.Println(output)

		err = saveToCSV("output.csv", output)
		if err != nil {
			log.Fatalf("Error saving to CSV: %v", err)
		}

	} else {
		fmt.Println("No channel statistics found.")
	}
}

func getChannelVideos(service *youtube.Service, channelID string) ([]*youtube.Video, error) {
	channelContent, err := service.Channels.List([]string{"contentDetails"}).Id(channelID).Do()
	if err != nil {
		return nil, err
	}

	if len(channelContent.Items) == 0 {
		return nil, fmt.Errorf("no channel content details found")
	}

	uploadsPlaylistID := channelContent.Items[0].ContentDetails.RelatedPlaylists.Uploads

	playlistItems, err := service.PlaylistItems.List([]string{"snippet"}).
		PlaylistId(uploadsPlaylistID).
		MaxResults(50).
		Do()
	if err != nil {
		return nil, err
	}

	if len(playlistItems.Items) == 0 {
		return nil, fmt.Errorf("no videos found in the uploads playlist")
	}

	var videos []*youtube.Video
	for _, item := range playlistItems.Items {
		videoID := item.Snippet.ResourceId.VideoId
		video, err := service.Videos.List([]string{"snippet", "contentDetails", "statistics"}).Id(videoID).Do()
		if err != nil {
			return nil, err
		}
		videos = append(videos, video.Items[0])
	}

	return videos, nil
}

type sortByFunc func(v1, v2 *youtube.Video) bool

func sortByViewsDesc(v1, v2 *youtube.Video) bool {
	return v1.Statistics.ViewCount > v2.Statistics.ViewCount
}

func sortByViewsAsc(v1, v2 *youtube.Video) bool {
	return v1.Statistics.ViewCount < v2.Statistics.ViewCount
}

func sortByLikesDesc(v1, v2 *youtube.Video) bool {
	return v1.Statistics.LikeCount > v2.Statistics.LikeCount
}

func sortByLikesAsc(v1, v2 *youtube.Video) bool {
	return v1.Statistics.LikeCount < v2.Statistics.LikeCount
}

func sortByCommentsDesc(v1, v2 *youtube.Video) bool {
	return v1.Statistics.CommentCount > v2.Statistics.CommentCount
}

func sortByCommentsAsc(v1, v2 *youtube.Video) bool {
	return v1.Statistics.CommentCount < v2.Statistics.CommentCount
}

func getTopVideosCSV(title string, videos []*youtube.Video, count int, sorter sortByFunc) string {
	var result strings.Builder
	result.WriteString(title + "\n")
	result.WriteString("Rank,Title,Views,Likes,Comments\n")

	sort.Slice(videos, func(i, j int) bool {
		return sorter(videos[i], videos[j])
	})

	for i, video := range videos[:count] {
		result.WriteString(fmt.Sprintf("%d,\"%s\",%d,%d,%d\n",
			i+1, video.Snippet.Title, video.Statistics.ViewCount, video.Statistics.LikeCount, video.Statistics.CommentCount))
	}

	return result.String()
}

func saveToCSV(filename, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}
