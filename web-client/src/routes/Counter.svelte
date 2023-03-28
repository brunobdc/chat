<script>
	import { browser } from '$app/environment';
	let messages = [];
	let sendMessage = ""
	let ws = null

	if (browser) {

		ws = new WebSocket("ws://localhost:12345/ws")

		ws.addEventListener("message", ev => {
			const data = JSON.parse(ev.data)
			messages.push(data)
			messages = messages
		})
	}

	const handleClick = () => {
		ws.send(sendMessage)
		sendMessage = ""
	}

</script>

<div class="messages">
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
</div>

<style>
	
</style>
