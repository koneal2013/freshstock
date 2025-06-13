import { useEffect, useState } from 'react'
import './App.css'

const API_URL = 'http://localhost:8080/api/v1/produce'

function App() {
  const [produce, setProduce] = useState([])
  const [search, setSearch] = useState('')
  const [form, setForm] = useState({ code: '', name: '', unit_price: '' })
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  // Fetch produce list (optionally with search)
  const fetchProduce = async (query = '') => {
    setLoading(true)
    setError('')
    try {
      const url = query ? `${API_URL}/?q=${encodeURIComponent(query)}` : `${API_URL}/`
      const res = await fetch(url)
      if (!res.ok) throw new Error('Failed to fetch produce')
      const data = await res.json()
      // Handle both array and single object responses
      let items = Array.isArray(data) ? data : [data]
      // Normalize unit_price to always be a string for display
      items = items.map(item => ({
        ...item,
        unit_price: typeof item.unit_price === 'string' ? item.unit_price : `$${Number(item.unit_price).toFixed(2)}`
      }))
      setProduce(items)
    } catch (e) {
      setProduce([])
      setError(e.message)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchProduce()
  }, [])

  // Handle form input
  const handleChange = e => {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  // Add produce
  const handleAdd = async e => {
    e.preventDefault()
    setError('')
    try {
      const res = await fetch(`${API_URL}/`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          code: form.code,
          name: form.name,
          unit_price: parseFloat(form.unit_price)
        })
      })
      if (!res.ok) {
        let errMsg = 'Failed to add produce'
        try {
          const err = await res.json()
          errMsg = err.error || errMsg
        } catch {
          // Intentionally left blank: error message will use default
        }
        throw new Error(errMsg)
      }
      setForm({ code: '', name: '', unit_price: '' })
      fetchProduce()
    } catch (e) {
      setError(e.message)
    }
  }

  // Delete produce
  const handleDelete = async code => {
    setError('')
    try {
      const res = await fetch(`${API_URL}/${code}`, { method: 'DELETE' })
      if (!res.ok && res.status !== 404) {
        const err = await res.json()
        throw new Error(err.error || 'Failed to delete produce')
      }
      fetchProduce(search)
    } catch (e) {
      setError(e.message)
    }
  }

  // Search produce
  const handleSearch = e => {
    e.preventDefault()
    fetchProduce(search)
  }

  return (
    <div className="container">
      <h1>FreshStock Produce Inventory</h1>
      <form onSubmit={handleSearch} style={{ marginBottom: 24 }}>
        <input
          type="text"
          placeholder="Search produce by name..."
          value={search}
          onChange={e => setSearch(e.target.value)}
          style={{ padding: 8, width: 240, marginRight: 8 }}
        />
        <button type="submit">Search</button>
        <button type="button" onClick={() => { setSearch(''); fetchProduce('') }} style={{ marginLeft: 8 }}>Clear</button>
      </form>

      <form onSubmit={handleAdd} className="add-form">
        <h2>Add Produce</h2>
        <input
          name="code"
          type="text"
          placeholder="Code (19 chars, e.g. A12T-4GH7-QPL9-3N4M)"
          value={form.code}
          onChange={handleChange}
          required
          minLength={19}
          maxLength={19}
        />
        <input
          name="name"
          type="text"
          placeholder="Name"
          value={form.name}
          onChange={handleChange}
          required
        />
        <input
          name="unit_price"
          type="number"
          step="0.01"
          min="0"
          placeholder="Unit Price"
          value={form.unit_price}
          onChange={handleChange}
          required
        />
        <button type="submit">Add</button>
      </form>

      {error && <div style={{ color: 'red', margin: 12 }}>{error}</div>}
      {loading ? <p>Loading...</p> : (
        <table className="produce-table">
          <thead>
            <tr>
              <th>Code</th>
              <th>Name</th>
              <th>Unit Price</th>
              <th>Action</th>
            </tr>
          </thead>
          <tbody>
            {produce.length === 0 ? (
              <tr><td colSpan={4}>No produce found.</td></tr>
            ) : produce.map(item => (
              <tr key={item.code}>
                <td>{item.code}</td>
                <td>{item.name}</td>
                <td>{item.unit_price}</td>
                <td>
                  <button onClick={() => handleDelete(item.code)} style={{ color: 'red' }}>Delete</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  )
}

export default App
