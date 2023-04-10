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

<main class="flex flex-col justify-around items-center h-[90vh] m-8">
  <MessageBox messages={messages} />
  <MessageForm on:submit={(event) => { ws.send(event.detail.message) }} />
</main>

