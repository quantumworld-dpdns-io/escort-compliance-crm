document.addEventListener('alpine:init', () => {
  Alpine.store('auth', {
    user: null,
    token: null,
    isAuthenticated: false,

    init() {
      const token = localStorage.getItem('token');
      const user = localStorage.getItem('user');
      if (token && user) {
        this.token = token;
        this.user = JSON.parse(user);
        this.isAuthenticated = true;
      }
    },

    login(token, user) {
      this.token = token;
      this.user = user;
      this.isAuthenticated = true;
      localStorage.setItem('token', token);
      localStorage.setItem('user', JSON.stringify(user));
    },

    logout() {
      this.token = null;
      this.user = null;
      this.isAuthenticated = false;
      localStorage.removeItem('token');
      localStorage.removeItem('user');
    },
  });
});
