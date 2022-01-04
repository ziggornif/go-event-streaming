import Tweet from './tweet.js';
import { dateFormatOpts } from "./utils.js"

/**
 * Append tweets in the tweets DOM element
 * @param domTweets
 * @param tweets
 */
function appendTweets(domTweets, tweets) {
  (tweets || []).forEach(tweet => {
    const elem = document.createElement("div");
    const header = document.createElement("div");
    const eventDate = new Intl.DateTimeFormat("fr-FR", dateFormatOpts).format(
      new Date(tweet.date)
    );
    header.innerText = `${tweet.author} - ${eventDate}`;

    const message = document.createElement("div");
    message.innerText = tweet.message

    const footer = document.createElement("div");
    const likesSpan = document.createElement("span")
    likesSpan.innerHTML = `${tweet.likes} ❤️`
    footer.appendChild(likesSpan);
    const likeBtn = document.createElement("button");
    likeBtn.id = `${tweet.id}-likebtn`;
    likeBtn.textContent = `Like ❤️`;
    likeBtn.addEventListener("click", function () {
      tweet.increase();
      likesSpan.innerHTML = `${tweet.likes} ❤️`
    })
    footer.classList.add(`${tweet.id}-footer`);
    footer.appendChild(likeBtn);

    elem.appendChild(header);
    elem.appendChild(message);
    elem.appendChild(footer);

    elem.classList.add("event");
    domTweets.insertBefore(elem, domTweets.firstChild);
  });
}

/**
 * Get tweets
 * @returns {Promise<Array<Tweet>>}
 */
async function getTweets() {
  const response = await fetch("http://localhost:8080/tweets")
  const jsonResp = await response.json();

  const tweets = [];

  for(const item of jsonResp) {
    tweets.push(new Tweet({
      id: item.id,
      author: item.author,
      message: item.message,
      date: item.created_at,
      likes: item.likes,
    }))
  }

  return tweets;
}

/**
 * Create tweet
 * @returns {Promise<any>}
 */
async function postTweet() {
  const message = document.getElementById('message').value;
  const author = document.getElementById('author').value;

  if(!message || !author) return

  const resp = await fetch("/tweets", {
    method: "POST",
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      message,
      author
    })
  })

  document.getElementById('message').value = null;
  document.getElementById('author').value = null;
  const respBody = await resp.json()

  const created = new Tweet({
    id: respBody.id,
    author: respBody.author,
    message: respBody.message,
    date: respBody.created_at,
    likes: respBody.likes,
  })
  return created;
}

function likeTweet(tweetId) {
  if (!tweetId) return

  return fetch(`/tweets/${tweetId}/likes`, {
    method: "POST",
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
  })
}

export {
  getTweets,
  appendTweets,
  likeTweet,
  postTweet,
}