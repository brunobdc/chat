<script>
    export let messages = [];
</script>

<div class="message-box">
{#each messages as msg}
    <p>{msg.sender}: {msg.content}</p>
{/each}
</div>

<style>
    .message-box {
        width: 50vw;
        height: 75%;
        border: 2px solid black;
        border-radius: 5px;
    }
</style>