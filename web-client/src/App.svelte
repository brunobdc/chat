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

<main>
  <div class="message-box">
    {#each messages as msg}
        <p>{msg.sender}: {msg.content}</p>
    {/each}
  </div>
  <div>
      <input type="text" bind:value={sendMessage} />
      <button on:click={handleClick}>SEND</button>
  </div>
</main>

<style>
  main {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    align-content: space-between;
    height: 85vh;
  }

  .message-box {
    width: 50vw;
    height: 100%;
    margin-top: 5vh;
  }
</style>