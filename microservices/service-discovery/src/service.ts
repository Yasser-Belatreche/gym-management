import { RedisClient } from './redis';

class ServiceDiscovery {
    private readonly KEY = 'services';

    constructor(private readonly client: RedisClient) {}

    async register(name: string, url: string): Promise<Service> {
        name = name.toLowerCase().trim();
        if (!name) {
            throw new Error('Invalid name');
        }

        const u = new URL(url);
        if (!u.protocol.startsWith('http')) {
            throw new Error('Invalid URL');
        }

        const service: Service = {
            name,
            url: u.protocol + '//' + u.host,
        };

        await this.client.hSet(this.KEY, service.name, service.url);

        return service;
    }

    async getServices(): Promise<Service[]> {
        const services: Service[] = [];
        const data = await this.client.hGetAll(this.KEY);

        for (const name in data) {
            services.push({ name, url: data[name] });
        }

        return services;
    }

    async getService(name: string): Promise<Service | undefined> {
        const url = await this.client.hGet(this.KEY, name);

        if (!url) {
            return undefined;
        }

        return { name, url };
    }

    async deleteService(name: string): Promise<void> {
        await this.client.hDel(this.KEY, name);
    }
}

interface Service {
    name: string;
    url: string;
}

export { ServiceDiscovery };
