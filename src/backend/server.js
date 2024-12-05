import express, { json } from 'express'
import { Pool } from 'pg'
import cors from 'cors'

const app = express()
const pool = new Pool({
  user: 'postgres',
  host: 'localhost',
  database: 'postgres',
  password: '123',
  port: 5432,
})

app.use(cors())
app.use(json())

app.get('/leaderboard', async (req, res) => {
  try {
    const result = await pool.query('SELECT * FROM leaderboard ORDER BY score DESC LIMIT 10')
    res.json(result.rows)
  } catch (err) {
    res.status(500).send(err.message)
  }
})

app.listen(3000, () => console.log('Server running on port 3000'))
