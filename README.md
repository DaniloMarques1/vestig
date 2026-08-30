# Vestig

A minimal, terminal-first CLI habit tracker designed to help you build consistency and keep track of your daily routines right from your command line.

With **vestig**, you can register habits you want to track, log their completion, inspect streaks, and soft-delete habits you no longer need.

## Features & Usage

### 1. Add a Habit (`add`)

Register a new habit you wish to start tracking:

    vestig add "Play Dark Souls 3"

### 2. List Habits (`list`)

List tracked habits. By default, only active habits are displayed:

    vestig list

To view all habits (including inactive ones), pass the `--all` flag:

    vestig list --all

### 3. Log Habit Execution (`execute`)

Mark a habit as executed by providing its ID. This records an execution entry for today:

    vestig execute 2

*(where `2` is the ID of the habit, e.g., "Go to the gym")*

#### Custom Date Logging
If you forgot to log a habit on a previous day, use the `--date` flag to log it for a specific date:

    vestig execute 2 --date 2026-08-28

### 4. Remove a Habit (`delete`)

Deactivate or soft-delete a habit from your active list:

    vestig delete 1

### 5. View Habit Details & Streaks (`view`)

View detailed execution history, current streak stats, and a weekly activity grid for a specific habit:

    vestig view 1

**Output Example:**

    Study Go (#1) — 🔥 5-day streak

    Sat  Sun  Mon  Tue  Wed  Thu  Fri
     ✔    ✔    ✔    ✔    ✔    ✘    ✔

---

## 🚧 Pending Features / Roadmap

- [ ] **`edit` command**: Ability to rename existing habits.
