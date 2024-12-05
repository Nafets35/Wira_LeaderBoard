<template>
  <div class="leaderboard">
    <h1>Wira Player Leaderboard</h1>
    <div class="controls">
      <button @click="addRandomPlayer">Sort</button>
      <button @click="resetLeaderboard">Reset Leaderboard</button>
    </div>
    <table>
      <thead>
        <tr>
          <th>No.</th>
          <th>User Name</th>
          <th>Class Name</th>
          <th>Score</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(player, index) in sortedPlayers" :key="player.id">
          <td>{{ index + 1 }}</td>
          <td>{{ player.name }}</td>
          <td>{{ player.class }}</td>
          <td>{{ player.score }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
export default {
  data() {
    return {
      players: [
        { id: 1, name: 'Alice', class: 'Cleric', score: 1200 },
        { id: 2, name: 'Bob', class: 'Archer', score: 950 },
        { id: 3, name: 'Charlie', class: 'Mage', score: 700 },
      ],
    }
  },
  computed: {
    sortedPlayers() {
      return [...this.players].sort((a, b) => b.score - a.score)
    },
  },
  methods: {
    addRandomPlayer() {
      const newPlayer = {
        id: Date.now(),
        name: `Player ${this.players.length + 1}`,
        class: `Fighter ${this.players.length + 1}`,
        score: Math.floor(Math.random() * 1500),
      }
      this.players.push(newPlayer)
    },
    resetLeaderboard() {
      this.players = []
    },
  },
}
</script>

<style>
.leaderboard {
  width: 100%;
  text-align: center;
}

table {
  width: 100%;
  border-collapse: collapse;
  table-layout: auto;
}

th,
td {
  padding: 10px;
  border: 1px solid #ccc;
}

thead {
  background-color: #f4f4f4;
}

.controls {
  margin-top: 20px;
}

button {
  padding: 10px 15px;
  margin: 5px;
  cursor: pointer;
  border: none;
  background-color: #007bff;
  color: white;
  border-radius: 5px;
}

button:hover {
  background-color: #0056b3;
}
</style>
