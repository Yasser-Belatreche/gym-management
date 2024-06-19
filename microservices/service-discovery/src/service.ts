import * as crypto from 'crypto';

import { RedisClient } from './redis.js';

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
            id: crypto.randomUUID(),
            name,
            url: u.protocol + '//' + u.host,
        };

        await this.client.hSet(this.KEY, service.id, JSON.stringify(service));

        return service;
    }

    async updateService(id: string, name: string, url: string): Promise<{ id: string }> {
        const u = new URL(url);
        if (!u.protocol.startsWith('http')) {
            throw new Error('Invalid URL');
        }

        const exists = await this.client.hExists(this.KEY, id);
        if (!exists) {
            throw new Error('Service not found');
        }

        const service: Service = {
            id,
            name,
            url: u.protocol + '//' + u.host,
        };

        await this.client.hSet(this.KEY, service.id, JSON.stringify(service));

        return { id };
    }

    async getServices(): Promise<Service[]> {
        const services: Service[] = [];
        const data = await this.client.hGetAll(this.KEY);

        for (const id in data) {
            const service = JSON.parse(data[id]) as Service;

            services.push(service);
        }

        return services;
    }

    async getService(id: string): Promise<Service | undefined> {
        const data = await this.client.hGet(this.KEY, id);

        if (!data) {
            return undefined;
        }

        return JSON.parse(data) as Service;
    }

    async deleteService(id: string): Promise<void> {
        const exists = await this.client.hExists(this.KEY, id);
        if (!exists) {
            throw new Error('Service not found');
        }

        await this.client.hDel(this.KEY, id);
    }
}

interface Service {
    id: string;
    name: string;
    url: string;
}

export { ServiceDiscovery };
