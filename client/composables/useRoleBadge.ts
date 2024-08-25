import type { Database } from '~/types/database';
type Role = Database['public']['Enums']['role_type'];

export function useRoleBadge(role: Role | null) {
    if (!role) return '';
    switch (role) {
        case 'verified':
            return 'https://hhtvcfarcfoxzylxjfij.supabase.co/storage/v1/object/public/attachments/verified.png?t=2024-08-25T05%3A21%3A57.073Z';
        case 'dev':
            return 'https://hhtvcfarcfoxzylxjfij.supabase.co/storage/v1/object/public/attachments/dev.png';
        case 'unverified':
            return '';
    }
}