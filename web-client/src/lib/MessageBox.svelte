<script>
    export let messages = [];
</script>

<div class="w-[50vw] h-full border rounded border-solid border-black bg-amber-50">
{#each messages as msg}
    <p>{msg.sender}: {msg.content}</p>
{/each}
</div>

