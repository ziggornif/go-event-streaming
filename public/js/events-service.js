import { dateFormatOpts } from "./utils.js"

/**
 * Append events in the events DOM element
 * @param domEvents
 * @param events
 */
function appendEvents(domEvents, events) {
  (events || []).forEach(item => {
    const elem = document.createElement("div");
    const header = document.createElement("div");
    const eventDate = new Intl.DateTimeFormat("fr-FR", dateFormatOpts).format(
      new Date(item.date)
    )
    header.innerText = `⚡️ event : ${item.messageType} - ${eventDate}`

    const text = document.createElement("div");
    if(item.messageType === 'tweet_created') {
      text.innerText = `📝 message : ${item.message}`
    } else {
      text.innerText = `liked tweet ID : ${item.id}`
    }

    const footer = document.createElement("div")
    if(item.messageType === 'tweet_created') {
      footer.innerText = `👤️ author : ${item.author}`
      footer.classList.add("footer")
    }

    elem.appendChild(header);
    elem.appendChild(text);
    elem.appendChild(footer);

    elem.classList.add("event")
    domEvents.insertBefore(elem, domEvents.firstChild);
  })
}

/**
 * Listen events
 * @returns {Promise<Object>}
 */
async function getEvents() {
  const response = await fetch("http://localhost:8080/listener/events")
  const events = await response.json();
  return events;
}

export {
  getEvents,
  appendEvents,
}