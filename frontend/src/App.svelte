<script lang="ts">

  import Header from "./components/Header.svelte"
  import Form from "./components/Form.svelte"
  import Rows from "./components/Rows.svelte"

  // tabs
  let items = ["List", "Form"]
  let activeItem = "List"

  const handleTabChange = e => {
    activeItem = e.detail
  }

  const added = () => {
    activeItem = "List"
  }
  
</script>

<Header {activeItem} {items} on:changeTab={handleTabChange} />

<main class="container">

  {#if activeItem === "List"}
    <Rows />
  {:else if activeItem === "Form"}
    <Form on:add={added}/>
  {:else}
    <p>What?</p>
  {/if}

</main>

<style>

  

</style>
