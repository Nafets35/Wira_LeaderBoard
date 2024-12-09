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
        <a href="#" @click="filterByClass('Human')">Human</a>
        <a href="#" @click="filterByClass('Numah')">Numah</a>
        <a href="#" @click="filterByClass('Bomoh')">Bomoh</a>
        <a href="#" @click="filterByClass('Dukun')">Dukun</a>
        <a href="#" @click="filterByClass('Hassasin')">Hassassin</a>
        <a href="#" @click="filterByClass('Asurak')">Asurak</a>
        <a href="#" @click="filterByClass('Bakunawan')">Bakunawan</a>
        <a href="#" @click="filterByClass('Orang Kayangan')">Orang Kayangan</a>
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
import "@/assets/styles/leaderboard.css";
import { defineComponent } from 'vue'
import type { Player } from '@/types/Player.ts'; // Adjust path if necessary

export default defineComponent({
  data() {
    return {
      players: [] as Player[],
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
