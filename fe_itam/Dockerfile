FROM node:18-alpine
RUN npm install -g pnpm
WORKDIR /app
COPY package.json ./
RUN pnpm setup
ENV PNPM_HOME="/root/.local/share/pnpm"
ENV PATH="${PATH}:${PNPM_HOME}"
RUN pnpm install 
RUN pnpm add -g nuxt@latest
RUN pnpm add nuxt@latest
COPY . .
RUN pnpm run build
EXPOSE 3000
CMD ["pnpm", "run", "start"]