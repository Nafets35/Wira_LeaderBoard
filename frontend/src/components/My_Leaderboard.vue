<template>
  <div class="leaderboard">
    <h1>Wira Player Leaderboard</h1>
    <div class="search-container">
      <input type="text" placeholder="Search by name..." v-model="searchQuery" class="search-bar" />
    </div>
    <div class="dropdown">
      <!-- Sort by dropdown -->
      <button class="dropbtn" @click.stop="toggleSortDropdown">Sort by:</button>
      <div v-if="isSortDropdownOpen" class="dropdown-content">
        <a href="#" @click="sort('score')">Score</a>
        <a href="#" @click="sort('name')">Name</a>
      </div>

      <!-- Class filter dropdown -->
      <button class="dropbtn" @click.stop="toggleClassDropdown">Class:</button>
      <div v-if="isClassDropdownOpen" class="dropdown-content">
        <a href="#" @click="filterByClass('All')">All</a>
        <a href="#" @click="filterByClass('Archer')">Archer</a>
        <a href="#" @click="filterByClass('Bard')">Bard</a>
        <a href="#" @click="filterByClass('Cleric')">Cleric</a>
        <a href="#" @click="filterByClass('Druid')">Druid</a>
        <!-- Show all classes -->
      </div>
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
        <tr v-for="(player, index) in paginatedPlayers" :key="player.id">
          <td>{{ index + 1 + (currentPage - 1) * itemsPerPage }}</td>
          <td>{{ player.name }}</td>
          <td>{{ player.class }}</td>
          <td>{{ player.score }}</td>
        </tr>
      </tbody>
    </table>

    <!-- Pagination Controls -->
    <div class="pagination">
      <button @click="previousPage" :disabled="currentPage === 1">Previous</button>
      <span>Page {{ currentPage }} of {{ totalPages }}</span>
      <button @click="nextPage" :disabled="currentPage === totalPages">Next</button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

interface Player {
  id: number
  name: string
  class: string
  score: number
}
export default defineComponent({
  data() {
    return {
      players: [] as Player[], // Explicitly type the players array
      sortBy: 'score',
      sortOrder: 'desc',
      isSortDropdownOpen: false,
      isClassDropdownOpen: false,
      filteredClass: null as string | null,
      currentPage: 1,
      itemsPerPage: 10,
      searchQuery: '',
    }
  },
  computed: {
    filteredPlayers(): Player[] {
      const query = this.searchQuery.trim().toLowerCase()
      let players = this.players

      if (this.filteredClass) {
        players = players.filter((player) => player.class === this.filteredClass)
      }

      if (query) {
        players = players.filter((player) => player.name.toLowerCase().includes(query))
      }

      return players
    },
    sortedPlayers(): Player[] {
      return [...this.filteredPlayers].sort((a, b) => {
        if (this.sortBy === 'score') {
          return this.sortOrder === 'desc' ? b.score - a.score : a.score - b.score
        } else if (this.sortBy === 'name') {
          return this.sortOrder === 'asc'
            ? a.name.localeCompare(b.name)
            : b.name.localeCompare(a.name)
        }
        return 0
      })
    },
    paginatedPlayers(): Player[] {
      const start = (this.currentPage - 1) * this.itemsPerPage
      const end = start + this.itemsPerPage
      return this.sortedPlayers.slice(start, end)
    },
    totalPages(): number {
      return Math.ceil(this.sortedPlayers.length / this.itemsPerPage)
    },
  },
  methods: {
    toggleSortDropdown() {
      this.isSortDropdownOpen = !this.isSortDropdownOpen
      this.isClassDropdownOpen = false
    },
    toggleClassDropdown() {
      this.isClassDropdownOpen = !this.isClassDropdownOpen
      this.isSortDropdownOpen = false
    },
    sort(choice: string) {
      if (this.sortBy === choice) {
        this.sortOrder = this.sortOrder === 'asc' ? 'desc' : 'asc'
      } else {
        this.sortBy = choice
        this.sortOrder = choice === 'name' ? 'asc' : 'desc'
      }
      this.isSortDropdownOpen = false
    },
    filterByClass(className: string) {
      if (className === 'All') {
        this.filteredClass = null
      } else {
        this.filteredClass = className
      }
      this.isClassDropdownOpen = false
    },
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++
      }
    },
    previousPage() {
      if (this.currentPage > 1) {
        this.currentPage--
      }
    },
    goToPage(page: number) {
      if (page >= 1 && page <= this.totalPages) {
        this.currentPage = page
      }
    },
    async fetchPlayers() {
      try {
        const response = await fetch('http://localhost:8080/leaderboard')
        if (!response.ok) {
          throw new Error('Failed to fetch leaderboard data')
        }
        const data: Player[] = await response.json()
        this.players = data
      } catch (error) {
        console.error('Error fetching leaderboard data:', error)
      }
    },
  },
  created() {
    this.fetchPlayers()
  },
})
</script>

<style>
.search-container {
  display: flex;
  justify-content: flex-start;
}

.search-bar {
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 5px;
  width: 200px;
}

.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
  gap: 10px;
}

.pagination button {
  padding: 5px 10px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.pagination button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

.dropdown {
  display: flex;
  justify-content: flex-end;
  gap: 20px;
}

.dropdown-content {
  display: none;
  position: absolute;
  background-color: #f9f9f9;
  min-width: 160px;
  box-shadow: 0px 8px 16px 0px rgba(0, 0, 0, 0.2);
  z-index: 1;
}

.dropdown-content a {
  color: black;
  padding: 12px 16px;
  text-decoration: none;
  display: block;
}

.dropdown-content a:hover {
  background-color: #f1f1f1;
}

.dropdown .dropdown-content {
  display: block;
}

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
