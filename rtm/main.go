package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/",home)
	http.HandleFunc("/events",eventHandler)
	fmt.Println("Server is listening on port 8000")
	http.ListenAndServe(":8000",nil)
}

func home(w  http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","text/html")
	w.Write([]byte("hello there"))
}

func eventHandler(w http.ResponseWriter, r *http.Request){
	// Generate unique client ID for this connection
	clientID := fmt.Sprintf("client-%d", time.Now().UnixNano()%10000)
	fmt.Printf("[%s] New client connected from %s\n", clientID, r.RemoteAddr)
	
	w.Header().Set("Content-Type","text/event-stream")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Connection","keep-alive")
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(200)
	// Check for Last-Event-ID header (for reconnection)
	lastEventID := r.Header.Get("Last-Event-ID")
	startIndex := 0
	
	tokens := []string{
		"Once", " upon", " a", " time,", " in", " a", " land", " far", " far", " away,", 
		" there", " lived", " a", " young", " programmer", " named", " Alex.", " Alex", 
		" had", " always", " been", " fascinated", " by", " the", " magic", " of", " code", 
		" and", " the", " endless", " possibilities", " it", " offered.", " Every", " day,", 
		" Alex", " would", " sit", " at", " the", " computer,", " crafting", " elegant", 
		" solutions", " to", " complex", " problems.", " The", " keyboard", " was", " like", 
		" a", " wand,", " and", " each", " line", " of", " code", " was", " a", " spell", 
		" that", " brought", " ideas", " to", " life.", " From", " simple", " websites", 
		" to", " complex", " applications,", " Alex's", " creations", " touched", " the", 
		" lives", " of", " thousands", " of", " people", " around", " the", " world.", 
		" And", " so", " the", " journey", " of", " discovery", " and", " innovation", 
		" continued,", " one", " line", " of", " code", " at", " a", " time...",
	}
	
	// If reconnecting, find where to resume
	if lastEventID != "" {
		fmt.Printf("[%s] Client reconnecting from event ID: %s\n", clientID, lastEventID)
		for i := range tokens {
			if fmt.Sprintf("msg-%d", i) == lastEventID {
				startIndex = i + 1
				break
			}
		}
	}
	
	// Set retry interval (3 seconds)
	w.Write([]byte("retry: 3000\n\n"))
	w.(http.Flusher).Flush()
	
	fmt.Printf("[%s] Starting stream from index %d\n", clientID, startIndex)
	
	// Stream tokens starting from the right position
	for i := startIndex; i < len(tokens); i++ {
		token := tokens[i]
		eventID := fmt.Sprintf("msg-%d", i)
		
		// Send message with ID
		content := fmt.Sprintf("id: %s\ndata: %s\n\n", eventID, token)
		w.Write([]byte(content))
		w.(http.Flusher).Flush()

		// Log every 10th message to show concurrent activity
		if i%10 == 0 {
			fmt.Printf("[%s] Sent token %d: %s\n", clientID, i, token)
		}

		time.Sleep(time.Millisecond * 500)
		
		// Simulate random disconnection (10% chance)
		// if i > 2 && i < len(tokens)-2 && time.Now().UnixNano()%10 == 0 {
		// 	fmt.Printf("Simulating disconnect at token %d\n", i)
		// 	return // This will close the connection
		// }
	}
	
	// Send completion message
	fmt.Printf("[%s] Stream completed\n", clientID)
	w.Write([]byte("id: complete\ndata: Stream completed!\n\n"))
	w.(http.Flusher).Flush()
}