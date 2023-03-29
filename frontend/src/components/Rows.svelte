<script lang="ts">

import { HandleGetGPMs } from "../../wailsjs/go/main/App"

let rows: any = []

function init() {
  getRows()
}

init()

function getRows() {
  HandleGetGPMs().then((res) => rows = res.rows)
}

</script>

<figure>
  <table role="grid">

    <thead>
      <tr>
        <th scope="col">#</th>
        <th scope="col">Website</th>
        <th scope="col">Username</th>
        <th scope="col">Email</th>
        <th scope="col">Password</th>
        <th scope="col">Notes</th>
      </tr>
    </thead>

    <tbody>

      {#each rows as row}
        <tr>
          <td>{row.id}</td>
          <td>{row.website}</td>
          <td>{row.username}</td>
          <td>{row.email}</td>
          <td>{row.password}</td>
          <td>{row.notes}</td>
        </tr>
      {/each}

    </tbody>

  </table>
</figure>

<style>

</style>