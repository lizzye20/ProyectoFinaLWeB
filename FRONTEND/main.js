document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.getElementById('login-form');
    const registerForm = document.getElementById('register-form');
    const reservationForm = document.getElementById('reservation-form');

    if (loginForm) {
        loginForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const username = document.getElementById('username').value;
            const password = document.getElementById('password').value;

            const response = await fetch('/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ username, password }),
            });

            if (response.ok) {
                const data = await response.json();
                localStorage.setItem('token', data.token);
                window.location.href = 'reservas.html';
            } else {
                alert('Login failed');
            }
        });
    }

    if (registerForm) {
        registerForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const username = document.getElementById('username').value;
            const password = document.getElementById('password').value;
            const email = document.getElementById('email').value;
            const phone = document.getElementById('phone').value;

            const response = await fetch('/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ username, password, email, phone }),
            });

            if (response.ok) {
                alert('Registration successful');
                window.location.href = 'login.html';
            } else {
                alert('Registration failed');
            }
        });
    }

    if (reservationForm) {
        reservationForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const carID = document.getElementById('car-id').value;
            const extras = document.getElementById('extras').value;
            const token = localStorage.getItem('token');

            const response = await fetch('/reservations', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`,
                },
                body: JSON.stringify({ car_id: carID, extras }),
            });

            if (response.ok) {
                alert('Reservation successful');
                loadReservations();
            } else {
                alert('Reservation failed');
            }
        });
    }

    if (window.location.pathname.endsWith('reservas.html')) {
        loadReservations();
    }

    async function loadReservations() {
        const token = localStorage.getItem('token');

        const response = await fetch('/reservations', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`,
            },
        });

        if (response.ok) {
            const reservations = await response.json();
            const reservationsList = document.getElementById('reservations-list');
            reservationsList.innerHTML = '';

            reservations.forEach((reservation) => {
                const listItem = document.createElement('li');
                listItem.textContent = `Car ID: ${reservation.car_id}, Extras: ${reservation.extras}`;
                reservationsList.appendChild(listItem);
            });
        } else {
            alert('Failed to load reservations');
        }
    }
});
