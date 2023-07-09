import React from 'react';
import { Link } from 'react-router-dom';
import '../css/Dashboard.css';

const Dashboard = () => {
    return (
        <div className="dashboard-container">
            <div className="sidebar">
                <h2>Menu</h2>
                <ul className="menu-list">
                    <li>
                        <Link to="/dashboard/current-image">Current Image</Link>
                    </li>
                    <li>
                        <Link to="/dashboard/history-image">History Image</Link>
                    </li>
                </ul>
            </div>
            <div className="content">
                <h1>Dashboard</h1>
                {/* Konten halaman dashboard */}
            </div>
        </div>
    );
};

export default Dashboard;
