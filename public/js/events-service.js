import { dateFormatOpts } from "./utils.js"

/**
 * Append events in the events DOM element
 * @param domEvents
 * @param events
 */
function appendEvent(domEvents, event) {
  const elem = document.createElement("div");
  const header = document.createElement("div");
  const eventDate = new Intl.DateTimeFormat("fr-FR", dateFormatOpts).format(
    new Date(event.date)
  )
  header.innerText = `⚡️ event : ${event.messageType} - ${eventDate}`

  const text = document.createElement("div");
  if(event.messageType === 'tweet_created') {
    text.innerText = `📝 message : ${event.message}`
  } else {
    text.innerText = `liked tweet ID : ${event.id}`
  }

  const footer = document.createElement("div")
  if(event.messageType === 'tweet_created') {
    footer.innerText = `👤️ author : ${event.author}`
    footer.classList.add("footer")
  }

  elem.appendChild(header);
  elem.appendChild(text);
  elem.appendChild(footer);

  elem.classList.add("event")
  domEvents.insertBefore(elem, domEvents.firstChild);
}

export {
  appendEvent,
}