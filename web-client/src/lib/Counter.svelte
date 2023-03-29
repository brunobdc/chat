<script>
  let messages = [];
  let sendMessage = "";

  const ws = new WebSocket(`ws://${location.host}/ws`)

  ws.addEventListener("message", ev => {
    const data = JSON.parse(ev.data)
    messages.push(data)
    messages = messages
  })

  const handleClick = () => {
		ws.send(sendMessage)
		sendMessage = ""
	}
</script>

<div>
  {#each messages as msg}
    <p>{msg.sender}: {msg.content}</p>
  {/each}
</div>
<div>
  <input type="text" bind:value={sendMessage} />
  <button on:click={handleClick}>
    SEND
  </button>
</div>
