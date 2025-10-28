FROM node:18-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

# Собираем приложение
RUN npm run build

# Используем переменные из docker-compose
EXPOSE ${VITE_APP_PORT:-5173} ${FRONTEND_DEBUG_PORT:-9229}

CMD ["npm", "run", "dev"]