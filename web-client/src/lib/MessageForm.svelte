<script>
  import { createEventDispatcher } from "svelte";

    let message = "";
    const submit = createEventDispatcher();

    const onSubmit = () => {
        submit("submit", { message });
        message = "";
    }
</script>

<form on:submit|preventDefault={onSubmit}>
    <textarea rows="4" bind:value={message} />
    <button type="submit">SEND</button>
</form> 

<style>
    textarea {
        resize: none;
    }
</style>