import {getEvents, appendEvents} from "./events-service.js";
import {getTweets, appendTweets, postTweet, likeTweet} from "./tweets-service.js";

window.onload = async function () {
  const eventsDom = document.getElementById("events");
  const tweetDom = document.getElementById("tweets");

  const tweets = await getTweets();
  appendTweets(tweetDom, tweets)

  setInterval(async () => {
    const events = await getEvents();
    appendEvents(eventsDom, events)
  }, 1000)

  document.getElementById("tweetbtn").addEventListener("click", async function () {
    const tweet = await postTweet();
    appendTweets(tweetDom, [tweet])
  })
};