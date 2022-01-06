import {appendEvent} from "./events-service.js";
import {getTweets, appendTweets, postTweet, likeTweet} from "./tweets-service.js";

window.onload = async function () {
  const eventsDom = document.getElementById("events");
  const tweetDom = document.getElementById("tweets");

  const tweets = await getTweets();
  appendTweets(tweetDom, tweets)

  document.getElementById("tweetbtn").addEventListener("click", async function () {
    const tweet = await postTweet();
    appendTweets(tweetDom, [tweet])
  })

  const conn = new WebSocket("ws://localhost:8080/listener/ws");

  conn.onmessage = function (evt) {
    appendEvent(eventsDom, JSON.parse(evt.data));
  };
};