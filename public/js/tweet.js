import { likeTweet } from "./tweets-service.js"

class Tweet {
  constructor({id, author, message, date, likes}) {
    this.id = id;
    this.author = author;
    this.message = message;
    this.date = date;
    this.likes = likes;
  }

  increase() {
    this.likes += 1;
    return likeTweet(this.id);
  }
}

export default Tweet;