<script>
  import MessageBox from "./lib/MessageBox.svelte";
  import MessageForm from "./lib/MessageForm.svelte";

  let messages = [];

  const ws = new WebSocket(`ws://${location.host}/ws`)

  ws.addEventListener("message", ev => {
      const data = JSON.parse(ev.data)
      messages.push(data)
      messages = messages
  })
</script>

<main>
  <MessageBox messages={messages} />
  <MessageForm on:submit={(event) => { ws.send(event.detail.message) }} />
</main>

<style>
  main {
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;
    height: 90vh;
    margin: 2rem;
  }
</style>