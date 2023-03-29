<script lang="ts">

import { HandleInsertGPM } from "../../wailsjs/go/main/App"
import { createEventDispatcher } from "svelte"

let dispatch = createEventDispatcher()

let form

const handleFormSubmit = () => {
  const data = new FormData(form)
  const fj: any = Object.fromEntries(data.entries())

  HandleInsertGPM(fj).then(res => console.log(res)).then(() => dispatch("add"))
}

</script>

<form on:submit|preventDefault={handleFormSubmit} bind:this={form}>

  <label for="website">
    Website
    <input
      type="text"
      id="website"
      name="website"
      placeholder="Website"
      required >
  </label>

  <label for="username">
    Username
    <input
      type="text"
      id="username"
      name="username"
      placeholder="Username"
      required >
  </label>

  <label for="email">
    Email
    <input
      type="email"
      id="email"
      name="email"
      placeholder="Email"
      required >
  </label>

  <label for="password">
    Password
    <input
      type="text"
      id="password"
      name="password"
      placeholder="Password"
      required >
  </label>

  <label for="notes">
    Notes
    <textarea
      name="notes"
      id="notes" >
    </textarea>
  </label>

  <button type="submit" class="outline">Add</button>

</form>

<style>
form {
  max-width: 640px;
  margin-inline: auto;
}
</style>