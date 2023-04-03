<script lang="ts">

  import Header from "./components/Header.svelte"
  import Form from "./components/Form.svelte"
  import Rows from "./components/Rows.svelte"

  // tabs
  let items = ["List", "Form"]
  let activeItem = "List"

  let showModal = false
  let dialog

  const handleTabChange = e => {
    activeItem = e.detail
  }

  const added = () => {
    activeItem = "List"
  }

  const toggleModal = (what="nothing") => {
    showModal = !showModal
    showModal ? dialog.show() : dialog.close()
    console.log(`modal toggle: ${showModal} action: ${what}`)
  }
  
</script>

<Header {activeItem} {items} on:changeTab={handleTabChange} />

<main class="container">

  {#if activeItem === "List"}
    <Rows />

    <button
      class="contrast"
      data-target="modal-example"
      on:click={() => toggleModal()} >
        Toggle Modal
    </button>

    <!-- Modal -->
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <dialog
      bind:this={dialog}
      on:click|self={() => toggleModal()} >

      <article on:click|stopPropagation>
        <a href="#close"
          aria-label="Close"
          class="close"
          on:click|preventDefault={() => toggleModal()}>
        </a>
        <h3>Confirm your action!</h3>
        <p>
          Cras sit amet maximus risus. 
          Pellentesque sodales odio sit amet augue finibus pellentesque. 
          Nullam finibus risus non semper euismod.
        </p>
        <footer>
          <a href="#cancel"
            role="button"
            class="secondary"
            on:click|preventDefault={() => toggleModal()}>
            Cancel
          </a>
          <a href="#confirm"
            role="button"
            on:click|preventDefault={() => toggleModal("action")}>
            Confirm
          </a>
        </footer>
      </article>
    </dialog>
    
  {:else if activeItem === "Form"}
    <Form on:add={added}/>
  {:else}
    <p>What?</p>
  {/if}


</main>

<style>

  

</style>
