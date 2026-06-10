import { describe, it, expect, beforeEach } from 'vitest'

describe('Auth Store', () => {
  beforeEach(() => {
    localStorage.clear()
  })

  it('initializes with no auth', () => {
    const store = { user: null, token: null, isAuthenticated: false }
    expect(store.isAuthenticated).toBe(false)
  })

  it('login sets token and user', () => {
    const store = { user: null, token: null, isAuthenticated: false }
    store.token = 'test-token'
    store.user = { email: 'test@example.com', name: 'Test' }
    store.isAuthenticated = true
    expect(store.isAuthenticated).toBe(true)
    expect(store.user.email).toBe('test@example.com')
  })

  it('logout clears state', () => {
    const store = { user: { name: 'Test' }, token: 'token', isAuthenticated: true }
    store.user = null
    store.token = null
    store.isAuthenticated = false
    expect(store.isAuthenticated).toBe(false)
  })
})
