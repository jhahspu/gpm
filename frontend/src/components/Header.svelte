<script>

import { createEventDispatcher } from "svelte"

const dispatch = createEventDispatcher()

export let activeItem
export let items

</script>

<nav class="container">
  <ul>
    <li><strong>GPM</strong></li>
  </ul>
  <ul>
    {#each items as item}
      <!-- svelte-ignore a11y-click-events-have-key-events -->
      <li on:click={() => dispatch("changeTab", item)}>
        <span class:active={item === activeItem}>{item}</span>
      </li>
    {/each}
  </ul>
</nav>

<style>

  ul li span {
    color: var(--primary);
    cursor: pointer;
  }

  ul li span:hover {
    color: var(--primary-hover);
    border-bottom: 1px solid var(--primary-hover);
  }
  
  ul li span.active {
    border-bottom: 1px dashed var(--primary-focus);
  }

</style>