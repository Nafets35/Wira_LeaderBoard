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
export default {
  data() {
    const names = [
      'Alice',
      'Bob',
      'Charlie',
      'Jamie',
      'Mia',
      'Liam',
      'Noah',
      'Olivia',
      'Sophia',
      'James',
      'Emma',
      'Ethan',
      'Isabella',
      'Mason',
      'Ava',
      'Logan',
      'Lucas',
      'Charlotte',
      'Amelia',
      'Henry',
    ]
    const classes = [
      'Cleric',
      'Archer',
      'Druid',
      'Bard',
      'Warrior',
      'Mage',
      'Paladin',
      'Rogue',
      'Sorcerer',
      'Monk',
    ]

    // Generate additional fake player entries
    const generatePlayers = (count: number) => {
      const players = []
      for (let i = 1; i <= count; i++) {
        players.push({
          id: i + 7, // Start from the next ID
          name: names[Math.floor(Math.random() * names.length)],
          class: classes[Math.floor(Math.random() * classes.length)],
          score: Math.floor(Math.random() * 1000) + 100, // Random score between 100 and 1100
        })
      }
      return players
    }

    return {
      players: [
        { id: 1, name: 'Alice', class: 'Cleric', score: 1200 },
        { id: 2, name: 'Bob', class: 'Archer', score: 700 },
        { id: 3, name: 'Charlie', class: 'Druid', score: 950 },
        { id: 4, name: 'Jamie', class: 'Bard', score: 800 },
        { id: 5, name: 'Derrick', class: 'Archer', score: 600 },
        { id: 6, name: 'Eden', class: 'Druid', score: 1050 },
        { id: 7, name: 'Fatin', class: 'Bard', score: 700 },
        ...generatePlayers(50),
      ],
      sortBy: 'score',
      sortOrder: 'desc',
      isSortDropdownOpen: false,
      isClassDropdownOpen: false,
      filteredClass: null as string | null,
      currentPage: 1, // Tracks the current page
      itemsPerPage: 10, // Number of players per page
      searchQuery: '', // For search functionality
    }
  },
  computed: {
    filteredPlayers() {
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
    sortedPlayers() {
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
    paginatedPlayers() {
      const start = (this.currentPage - 1) * this.itemsPerPage
      const end = start + this.itemsPerPage
      return this.sortedPlayers.slice(start, end)
    },
    totalPages() {
      return Math.ceil(this.sortedPlayers.length / this.itemsPerPage)
    },
  },
  methods: {
    toggleSortDropdown() {
      this.isSortDropdownOpen = !this.isSortDropdownOpen
      this.isClassDropdownOpen = false // Close the other dropdown
    },
    toggleClassDropdown() {
      this.isClassDropdownOpen = !this.isClassDropdownOpen
      this.isSortDropdownOpen = false // Close the other dropdown
    },
    sort(choice: string) {
      if (this.sortBy === choice) {
        this.sortOrder = this.sortOrder === 'asc' ? 'desc' : 'asc'
      } else {
        this.sortBy = choice
        this.sortOrder = choice === 'name' ? 'asc' : 'desc'
      }
      console.log(`Sorting by: ${choice}, Order: ${this.sortOrder}`)
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
  },
}
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
