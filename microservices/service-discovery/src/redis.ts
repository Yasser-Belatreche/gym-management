import { createClient } from 'redis';

export type RedisClient = ReturnType<typeof createClient>;

interface Config {
    url: string;
}

class Redis {
    private static instance: Redis;
    private readonly client: ReturnType<typeof createClient>;

    static Instance(config: Config): Redis {
        if (this.instance) return this.instance;

        this.instance = new Redis(config);

        return this.instance;
    }

    private constructor(config: Config) {
        this.client = createClient({ url: config.url });
    }

    Client(): RedisClient {
        return this.client;
    }

    async connect(): Promise<void> {
        await this.client.connect();
    }

    async healthCheck(): Promise<{ status: 'UP' | 'DOWN' }> {
        try {
            await this.client.ping();

            return { status: 'UP' };
        } catch (error) {
            return { status: 'DOWN' };
        }
    }
}

export { Redis };
