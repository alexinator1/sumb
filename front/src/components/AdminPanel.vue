<template>
    <div class="min-h-screen bg-gray-50">
      <!-- Header -->
      <header class="bg-white shadow-sm border-b border-gray-200">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div class="flex justify-between items-center h-16">
            <!-- Left side -->
            <div class="flex items-center space-x-3">
              <div class="w-8 h-8 bg-blue-600 rounded-lg flex items-center justify-center">
                <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
              </div>
              <h1 class="text-2xl font-bold text-gray-900">Admin Panel</h1>
            </div>
  
            <!-- Right side -->
            <div class="flex items-center space-x-4">
              <button class="p-2 text-gray-400 hover:text-gray-500">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
                </svg>
              </button>
              <div class="w-8 h-8 bg-gray-300 rounded-full flex items-center justify-center">
                <span class="text-sm font-medium text-gray-600">A</span>
              </div>
            </div>
          </div>
        </div>
      </header>
  
      <div class="flex">
        <!-- Sidebar -->
        <aside class="w-64 bg-white shadow-sm border-r border-gray-200 min-h-[calc(100vh-4rem)]">
          <nav class="mt-8">
            <div class="px-4 space-y-2">
              <a v-for="link in links" :key="link.name" :href="link.href" 
                 class="flex items-center px-3 py-2 text-gray-600 rounded-lg hover:bg-gray-50 hover:text-gray-900 group transition-colors"
                 :class="{ 'bg-blue-50 text-blue-700 border-r-2 border-blue-700': link.current }">
                <component :is="link.icon" class="w-5 h-5 mr-3" />
                <span class="font-medium">{{ link.name }}</span>
              </a>
            </div>
          </nav>
        </aside>
  
        <!-- Main content -->
        <main class="flex-1 p-8">
          <div class="max-w-7xl mx-auto">
            <!-- Stats -->
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
              <div v-for="stat in stats" :key="stat.name" 
                   class="bg-white rounded-lg shadow-sm border border-gray-200 p-6">
                <div class="flex items-center">
                  <div :class="['p-3 rounded-lg', stat.bgColor]">
                    <component :is="stat.icon" :class="['w-6 h-6', stat.iconColor]" />
                  </div>
                  <div class="ml-4">
                    <p class="text-sm font-medium text-gray-600">{{ stat.name }}</p>
                    <p class="text-2xl font-bold text-gray-900">{{ stat.value }}</p>
                  </div>
                </div>
              </div>
            </div>
  
            <!-- Table -->
            <div class="bg-white rounded-lg shadow-sm border border-gray-200">
              <div class="px-6 py-4 border-b border-gray-200">
                <div class="flex justify-between items-center">
                  <h2 class="text-lg font-semibold text-gray-900">Users</h2>
                  <button class="bg-blue-600 text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-blue-700 transition-colors flex items-center">
                    <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                    </svg>
                    Add New
                  </button>
                </div>
              </div>
              <div class="overflow-x-auto">
                <table class="w-full">
                  <thead>
                    <tr class="bg-gray-50 border-b border-gray-200">
                      <th v-for="column in columns" :key="column.key" 
                          class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                        {{ column.label }}
                      </th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-gray-200">
                    <tr v-for="user in users" :key="user.id" class="hover:bg-gray-50 transition-colors">
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                        {{ user.id }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                        {{ user.name }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
                        {{ user.email }}
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800">
                          {{ user.role }}
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap">
                        <span :class="[
                          'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
                          user.status === 'Active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                        ]">
                          {{ user.status }}
                        </span>
                      </td>
                      <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                        <div class="flex space-x-2">
                          <button @click="editUser(user)" 
                                  class="text-blue-600 hover:text-blue-900 transition-colors">
                            Edit
                          </button>
                          <button @click="deleteUser(user)" 
                                  class="text-red-600 hover:text-red-900 transition-colors">
                            Delete
                          </button>
                        </div>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </main>
      </div>
    </div>
  </template>
  
  <script>
  // Иконки как компоненты
  const DashboardIcon = {
    template: `
      <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
      </svg>
    `
  }
  
  const UsersIcon = {
    template: `
      <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197m13.5-9a2.5 2.5 0 11-5 0 2.5 2.5 0 015 0z" />
      </svg>
    `
  }
  
  const ProductsIcon = {
    template: `
      <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
      </svg>
    `
  }
  
  const OrdersIcon = {
    template: `
      <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
      </svg>
    `
  }
  
  const SettingsIcon = {
    template: `
      <svg fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
      </svg>
    `
  }
  
  export default {
    name: 'AdminPanel',
    components: {
      DashboardIcon,
      UsersIcon,
      ProductsIcon,
      OrdersIcon,
      SettingsIcon
    },
    data() {
      return {
        links: [
          { name: 'Dashboard', href: '#', icon: 'DashboardIcon', current: true },
          { name: 'Users', href: '#', icon: 'UsersIcon', current: false },
          { name: 'Products', href: '#', icon: 'ProductsIcon', current: false },
          { name: 'Orders', href: '#', icon: 'OrdersIcon', current: false },
          { name: 'Settings', href: '#', icon: 'SettingsIcon', current: false }
        ],
        stats: [
          { 
            name: 'Total Users', 
            value: '1,234', 
            icon: UsersIcon,
            bgColor: 'bg-blue-100',
            iconColor: 'text-blue-600'
          },
          { 
            name: 'Orders', 
            value: '567', 
            icon: OrdersIcon,
            bgColor: 'bg-green-100',
            iconColor: 'text-green-600'
          },
          { 
            name: 'Revenue', 
            value: '$12,345', 
            icon: ProductsIcon,
            bgColor: 'bg-purple-100',
            iconColor: 'text-purple-600'
          }
        ],
        columns: [
          { key: 'id', label: 'ID' },
          { key: 'name', label: 'Name' },
          { key: 'email', label: 'Email' },
          { key: 'role', label: 'Role' },
          { key: 'status', label: 'Status' },
          { key: 'actions', label: 'Actions' }
        ],
        users: [
          { id: 1, name: 'Леха', email: 'john@example.com', role: 'Admin', status: 'Active' },
          { id: 2, name: 'Арина', email: 'jane@example.com', role: 'User', status: 'Active' },
          { id: 3, name: 'Луна', email: 'bob@example.com', role: 'Editor', status: 'Inactive' }
        ]
      }
    },
    methods: {
      editUser(user) {
        alert(`Edit user: ${user.name}`)
      },
      deleteUser(user) {
        if (confirm(`Delete user ${user.name}?`)) {
          this.users = this.users.filter(u => u.id !== user.id)
        }
      }
    }
  }
  </script>