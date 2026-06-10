document.addEventListener('alpine:init', () => {
  // Global app state
  Alpine.data('app', () => ({
    loading: false,
    currentPage: 'dashboard',
    sidebarOpen: true,

    init() {
      this.$store.auth.init();
      this.$on('navigate', (page) => this.navigate(page));
    },

    navigate(page) {
      this.currentPage = page;
    },
  }));

  // Login page
  Alpine.data('loginPage', () => ({
    email: '',
    password: '',
    error: '',
    submitting: false,

    async handleLogin() {
      this.submitting = true;
      this.error = '';
      try {
        const res = await fetch('/api/v1/auth/login', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ email: this.email, password: this.password }),
        });
        const data = await res.json();
        if (!res.ok) throw new Error(data.message || 'Login failed');
        Alpine.store('auth').login(data.token, data.user || { email: this.email, name: this.email.split('@')[0] });
        this.$root.__x.$data.currentPage = 'dashboard';
      } catch (e) {
        this.error = e.message;
      } finally {
        this.submitting = false;
      }
    },
  }));

  // Register page
  Alpine.data('registerPage', () => ({
    displayName: '',
    email: '',
    password: '',
    role: 'client',
    error: '',
    submitting: false,

    async handleRegister() {
      this.submitting = true;
      this.error = '';
      try {
        const res = await fetch('/api/v1/auth/register', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ name: this.displayName, email: this.email, password: this.password, role: this.role }),
        });
        if (!res.ok) {
          const data = await res.json();
          throw new Error(data.message || 'Registration failed');
        }
        this.$dispatch('navigate', 'login');
      } catch (e) {
        this.error = e.message;
      } finally {
        this.submitting = false;
      }
    },
  }));

  // Dashboard
  Alpine.data('dashboardPage', () => ({
    stats: { companions: 24, activeBookings: 8, complianceScore: 96 },
  }));

  // Companions
  Alpine.data('companionsPage', () => ({
    search: '',
    showCreateModal: false,
    companions: [
      { id: '1', name: 'Alice M.', jurisdiction: 'US-CA', rate: 150, verified: true },
      { id: '2', name: 'Bob K.', jurisdiction: 'US-NV', rate: 200, verified: true },
      { id: '3', name: 'Clara S.', jurisdiction: 'US-NY', rate: 175, verified: false },
    ],
  }));

  // Bookings
  Alpine.data('bookingsPage', () => ({
    bookings: [
      { id: '1042', companion: 'Alice M.', date: '2025-06-15 14:00', jurisdiction: 'US-CA', status: 'confirmed' },
      { id: '1041', companion: 'Bob K.', date: '2025-06-14 10:00', jurisdiction: 'US-NV', status: 'completed' },
      { id: '1040', companion: 'Clara S.', date: '2025-06-16 16:00', jurisdiction: 'US-NY', status: 'pending' },
    ],
    statusClass(status) {
      return { confirmed: 'bg-green-100 text-green-700', completed: 'bg-blue-100 text-blue-700', pending: 'bg-yellow-100 text-yellow-700', cancelled: 'bg-red-100 text-red-700' }[status] || 'bg-gray-100';
    },
  }));

  // Compliance
  Alpine.data('compliancePage', () => ({
    score: 96,
    jurisdictions: [
      { id: '1', name: 'US-CA (California)', status: 'Active', compliant: true },
      { id: '2', name: 'US-NV (Nevada)', status: 'Active', compliant: true },
      { id: '3', name: 'US-NY (New York)', status: 'Review', compliant: false },
    ],
    alerts: [
      { id: '1', message: 'NY regulation update pending', severity: 'warning' },
    ],
  }));

  // Quantum
  Alpine.data('quantumPage', () => ({
    pqcKeys: [
      { id: '1', algorithm: 'Kyber768', created: '2025-06-01' },
      { id: '2', algorithm: 'Dilithium3', created: '2025-06-01' },
      { id: '3', algorithm: 'Falcon512', created: '2025-06-05' },
    ],
  }));

  // Screening
  Alpine.data('screeningPage', () => ({}));

  // Credentials
  Alpine.data('credentialsPage', () => ({}));

  // Payments
  Alpine.data('paymentsPage', () => ({}));

  // Messages
  Alpine.data('messagesPage', () => ({}));

  // Settings
  Alpine.data('settingsPage', () => ({}));
});
