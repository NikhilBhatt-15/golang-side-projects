package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)
func main(){
	var urlList = []string{
		"https://google.com",
		"https://youtube.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://go.dev",
		"https://gophercises.com",
		"https://medium.com",
		"https://reddit.com",
		"https://linkedin.com",
		"https://twitter.com",
		"https://facebook.com",
		"https://instagram.com",
		"https://netlify.com",
		"https://vercel.com",
		"https://cloudflare.com",
		"https://openai.com",
		"https://developer.mozilla.org",
		"https://dev.to",
		"https://hashnode.com",
		"https://docker.com",
		"https://kubernetes.io",
		"https://aws.amazon.com",
		"https://azure.microsoft.com",
		"https://digitalocean.com",
		"https://heroku.com",
		"https://supabase.com",
		"https://firebase.google.com",
		"https://vitejs.dev",
		"https://nextjs.org",
		"https://astro.build",
		"https://deno.land",
		"https://npmjs.com",
		"https://pnpm.io",
		"https://bun.sh",
		"https://yarnpkg.com",
		"https://postman.com",
		"https://insomnia.rest",
		"https://fig.io",
		"https://eslint.org",
		"https://prettier.io",
		"https://typescriptlang.org",
		"https://nodejs.org",
		"https://python.org",
		"https://rust-lang.org",
		"https://ruby-lang.org",
		"https://react.dev",
		"https://angular.io",
		"https://vuejs.org",
		"https://svelte.dev",
		"https://tailwindcss.com",
		"https://getbootstrap.com",
		"https://w3schools.com",
		"https://freecodecamp.org",
		"https://codecademy.com",
		"https://coursera.org",
		"https://udemy.com",
		"https://edx.org",
		"https://pluralsight.com",
		"https://frontendmentor.io",
		"https://css-tricks.com",
		"https://canva.com",
		"https://figma.com",
		"https://dribbble.com",
		"https://behance.net",
		"https://gitlab.com",
		"https://bitbucket.org",
		"https://codesandbox.io",
		"https://replit.com",
		"https://glitch.com",
		"https://jsfiddle.net",
		"https://codepen.io",
		"https://jsonplaceholder.typicode.com",
		"https://reqres.in",
		"https://httpbin.org",
		"https://mockapi.io",
		"https://rgbacolorpicker.com",
		"https://newsapi.org",
		"https://rapidapi.com",
		"https://unsplash.com",
		"https://pexels.com",
		"https://images.unsplash.com",
		"https://cdn.jsdelivr.net",
		"https://cdnjs.com",
		"https://fontawesome.com",
		"https://fonts.google.com",
		"https://material.io",
		"https://www.quora.com",
		"https://chakra-ui.com",
		"https://shadcn.dev",
		"https://tailgrids.com",
		"https://flowbite.com",
		"https://vitejs.dev",
		"https://play.golang.org",
		"https://pkg.go.dev",
		"https://thealgorists.com",
		"https://csstricks.com",
		"https://awesome-go.com",
		"https://realpython.com",
		"https://towardsdatascience.com",
		"https://machinelearningmastery.com",
		"https://stayfinder.nikhilbhatt.tech",
		"https://portfolio.nikhilbhatt.tech",
	}
		
		checkAsync(urlList)
}

func checkSync(urlList []string){
	start := time.Now()	
	for _,url := range urlList{
		fetch(url)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func checkAsync(urlList []string){
	start := time.Now()
	ch := make(chan string)
	for _,url:= range urlList{
		go asyncfetch(url,ch)
	}

	for range urlList{
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func asyncfetch(url string,ch chan string){
	start:= time.Now()

	resp,err := http.Get(url)
	if err!=nil{
		ch <- fmt.Sprint(err)
		return
	}
	
	nbytes,err := io.Copy(io.Discard,resp.Body)
	resp.Body.Close()
	if err !=nil{
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func fetch(url string){
	start := time.Now()
	 resp,err := http.Get(url)
	 if err!=nil{
		fmt.Printf("err: %v\n", err)
		return
	 }
	 nbytes,err := io.Copy(io.Discard,resp.Body)
	 if err !=nil{
		fmt.Printf("while reading %s: %v", url, err)
		return
	}
	resp.Body.Close()
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs %7d %s\n", secs, nbytes, url)
}