# --------------- Build Stage ---------------

FROM oven/bun as builder

WORKDIR /app

# Copy depencies first for better caching
COPY package*.json bun.lock* ./

# Install dependencies with conditional lockfile handling
RUN bun install $([ -f bun.lock ] && echo "--frozen-lockfile" || echo "")

COPY . .

# Update Next.js config to enable output (uncomment any line containing 'output')
RUN sed -i 's|//.output|output|g' next.config.ts

# Build app
RUN bun --bun run build

# --------------- Runner Stage ---------------

FROM gcr.io/distroless/nodejs18 AS runner

USER nonroot

# Copy only necessary files from builder
COPY --from=builder --chown=nonroot:nonroot /app/.next/standalone ./
COPY --from=builder --chown=nonroot:nonroot /app/public ./public
COPY --from=builder --chown=nonroot:nonroot /app/.next/static ./.next/static

# Expose port
EXPOSE 3000

# Environment variables
ENV NEXT_TELEMETRY_DISABLED=1 \
    PORT=3000 \
    HOSTNAME="0.0.0.0" \
    NODE_ENV=production

# Start application
CMD ["server.js"]