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
  Alpine.data('screeningPage', () => ({
    records: [
      { id: 'scr-001-abc123', screening_type: 'background_check', provider: 'sterling', status: 'verified', risk_flags: [] },
      { id: 'scr-002-def456', screening_type: 'identity_verification', provider: 'jumio', status: 'verified', risk_flags: [] },
      { id: 'scr-003-ghi789', screening_type: 'criminal_check', provider: 'checkr', status: 'flagged', risk_flags: ['medium_risk'] },
      { id: 'scr-004-jkl012', screening_type: 'reference_check', provider: 'sterling', status: 'pending', risk_flags: [] },
    ],
    async initiateScreening() {
      try {
        const res = await fetch('/api/v1/screening', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ user_id: 'user-1', companion_id: 'comp-1', screening_type: 'background_check', provider: 'sterling' }),
        });
        const data = await res.json();
        if (data.screening) this.records.unshift(data.screening);
      } catch (e) { console.error(e); }
    },
    async processScreening(id) {
      try {
        const res = await fetch(`/api/v1/screening/${id}/process`, { method: 'POST' });
        const data = await res.json();
        if (data.screening) {
          const idx = this.records.findIndex(r => r.id === id);
          if (idx !== -1) this.records[idx] = data.screening;
        }
      } catch (e) { console.error(e); }
    },
  }));

  // Credentials
  Alpine.data('credentialsPage', () => ({
    credentials: [
      { id: 'cred-001-abc', credential_type: 'identity_verification', issuer: 'escort-crm', subject: 'user-1', status: 'active', issued_at: '2025-06-01' },
      { id: 'cred-002-def', credential_type: 'age_verification', issuer: 'escort-crm', subject: 'user-2', status: 'active', issued_at: '2025-06-05' },
      { id: 'cred-003-ghi', credential_type: 'background_check', issuer: 'sterling', subject: 'user-3', status: 'revoked', issued_at: '2025-05-15' },
    ],
    async issueCredential() {
      try {
        const res = await fetch('/api/v1/credentials', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ holder_id: 'user-1', credential_type: 'identity_verification', issuer: 'escort-crm', subject: 'user-1' }),
        });
        const data = await res.json();
        if (data.credential) this.credentials.unshift(data.credential);
      } catch (e) { console.error(e); }
    },
    async verifyCredential(id) {
      try {
        const res = await fetch(`/api/v1/credentials/${id}/verify`);
        const data = await res.json();
        alert(data.valid ? 'Credential is valid' : 'Credential is invalid');
      } catch (e) { console.error(e); }
    },
    async revokeCredential(id) {
      if (!confirm('Are you sure you want to revoke this credential?')) return;
      try {
        await fetch(`/api/v1/credentials/${id}/revoke`, { method: 'POST' });
        const idx = this.credentials.findIndex(c => c.id === id);
        if (idx !== -1) this.credentials[idx].status = 'revoked';
      } catch (e) { console.error(e); }
    },
  }));

  // Payments
  Alpine.data('paymentsPage', () => ({
    payments: [
      { id: 'pay-001-abc', booking_id: 'book-001', amount: 150.00, method: 'card', status: 'completed', currency: 'USD' },
      { id: 'pay-002-def', booking_id: 'book-002', amount: 200.00, method: 'card', status: 'pending', currency: 'USD' },
      { id: 'pay-003-ghi', booking_id: 'book-003', amount: 75.00, method: 'bank_transfer', status: 'refunded', currency: 'USD' },
    ],
    async createPayment() {
      try {
        const res = await fetch('/api/v1/payments', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ booking_id: 'book-new', payer_id: 'user-1', payee_id: 'comp-1', amount: 100.00, method: 'card' }),
        });
        const data = await res.json();
        if (data.payment) this.payments.unshift(data.payment);
      } catch (e) { console.error(e); }
    },
    async processPayment(id) {
      try {
        const res = await fetch(`/api/v1/payments/${id}/process`, { method: 'POST' });
        const data = await res.json();
        if (data.payment) {
          const idx = this.payments.findIndex(p => p.id === id);
          if (idx !== -1) this.payments[idx] = { ...this.payments[idx], ...data.payment };
        }
      } catch (e) { console.error(e); }
    },
    async refundPayment(id) {
      const payment = this.payments.find(p => p.id === id);
      if (!payment) return;
      if (!confirm(`Refund $${payment.amount}?`)) return;
      try {
        await fetch(`/api/v1/payments/${id}/refund`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ amount: payment.amount }),
        });
        const idx = this.payments.findIndex(p => p.id === id);
        if (idx !== -1) this.payments[idx].status = 'refunded';
      } catch (e) { console.error(e); }
    },
  }));

  // Messages
  Alpine.data('messagesPage', () => ({
    searchQuery: '',
    selectedConversation: null,
    newMessage: '',
    conversations: [
      { id: 'conv-1', name: 'Alice M.', lastMessage: 'See you tomorrow at 2pm', lastMessageTime: '2m ago', unread: 2 },
      { id: 'conv-2', name: 'Bob K.', lastMessage: 'Thanks for the update', lastMessageTime: '15m ago', unread: 0 },
      { id: 'conv-3', name: 'Clara S.', lastMessage: 'Is the booking confirmed?', lastMessageTime: '1h ago', unread: 1 },
    ],
    messagesByConversation: {
      'conv-1': [
        { id: 'm1', sender_id: 'other', content: 'Hi! Are you available tomorrow?', time: '10:30 AM' },
        { id: 'm2', sender_id: 'me', content: 'Yes, I am! What time works for you?', time: '10:32 AM' },
        { id: 'm3', sender_id: 'other', content: 'See you tomorrow at 2pm', time: '10:35 AM' },
      ],
      'conv-2': [
        { id: 'm4', sender_id: 'me', content: 'The compliance check passed', time: '9:00 AM' },
        { id: 'm5', sender_id: 'other', content: 'Thanks for the update', time: '9:15 AM' },
      ],
      'conv-3': [
        { id: 'm6', sender_id: 'other', content: 'Is the booking confirmed?', time: '8:00 AM' },
      ],
    },
    get filteredConversations() {
      if (!this.searchQuery) return this.conversations;
      return this.conversations.filter(c => c.name.toLowerCase().includes(this.searchQuery.toLowerCase()));
    },
    get currentMessages() {
      return this.messagesByConversation[this.selectedConversation] || [];
    },
    selectConversation(id) {
      this.selectedConversation = id;
    },
    getCurrentConversation() {
      return this.conversations.find(c => c.id === this.selectedConversation);
    },
    async sendMessage() {
      if (!this.newMessage.trim() || !this.selectedConversation) return;
      const msg = {
        id: 'm-' + Date.now(),
        sender_id: 'me',
        content: this.newMessage,
        time: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' }),
      };
      if (!this.messagesByConversation[this.selectedConversation]) {
        this.messagesByConversation[this.selectedConversation] = [];
      }
      this.messagesByConversation[this.selectedConversation].push(msg);
      this.newMessage = '';
      try {
        await fetch('/api/v1/messages', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ conversation_id: this.selectedConversation, sender_id: 'me', recipient_id: 'other', content: msg.content }),
        });
      } catch (e) { console.error(e); }
    },
  }));

  // Settings
  Alpine.data('settingsPage', () => ({}));
});
