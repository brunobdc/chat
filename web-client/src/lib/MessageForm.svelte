<script>
  import { createEventDispatcher } from "svelte";

    let message = "";
    const submit = createEventDispatcher();

    const onSubmit = () => {
        submit("submit", { message });
        message = "";
    }
</script>

<form on:submit|preventDefault={onSubmit} class="flex items-end justify-between w-[50vw] gap-x-2 bg-slate-200 p-2">
    <span contenteditable role="textbox" class="w-full min-h-[1.5rem] max-h-20 overflow-y-auto p-2 bg-white rounded border border-black box-content" bind:innerHTML={message}/>
    <button type="submit" class="h-6 p-2 bg-white rounded border border-black text-center align-middle box-content">SEND</button>
</form> 

