export type Json =
  | string
  | number
  | boolean
  | null
  | { [key: string]: Json | undefined }
  | Json[]

export type Database = {
  public: {
    Tables: {
      posts: {
        Row: {
          attachment_url: string | null
          content: string
          created_at: string
          id: string
          parent_id: string | null
          profile_id: string | null
        }
        Insert: {
          attachment_url?: string | null
          content: string
          created_at?: string
          id?: string
          parent_id?: string | null
          profile_id?: string | null
        }
        Update: {
          attachment_url?: string | null
          content?: string
          created_at?: string
          id?: string
          parent_id?: string | null
          profile_id?: string | null
        }
        Relationships: [
          {
            foreignKeyName: "posts_parent_id_fkey"
            columns: ["parent_id"]
            isOneToOne: false
            referencedRelation: "posts"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "posts_parent_id_fkey"
            columns: ["parent_id"]
            isOneToOne: false
            referencedRelation: "posts_view"
            referencedColumns: ["post_id"]
          },
          {
            foreignKeyName: "posts_profile_id_fkey"
            columns: ["profile_id"]
            isOneToOne: false
            referencedRelation: "posts_view"
            referencedColumns: ["user_id"]
          },
          {
            foreignKeyName: "posts_profile_id_fkey"
            columns: ["profile_id"]
            isOneToOne: false
            referencedRelation: "profiles"
            referencedColumns: ["id"]
          },
        ]
      }
      profiles: {
        Row: {
          avatar_url: string
          bio: string | null
          created_at: string
          id: string
          role: Database["public"]["Enums"]["role_type"]
          username: string
        }
        Insert: {
          avatar_url?: string
          bio?: string | null
          created_at?: string
          id?: string
          role?: Database["public"]["Enums"]["role_type"]
          username: string
        }
        Update: {
          avatar_url?: string
          bio?: string | null
          created_at?: string
          id?: string
          role?: Database["public"]["Enums"]["role_type"]
          username?: string
        }
        Relationships: [
          {
            foreignKeyName: "profiles_id_fkey"
            columns: ["id"]
            isOneToOne: true
            referencedRelation: "users"
            referencedColumns: ["id"]
          },
        ]
      }
      reactions: {
        Row: {
          created_at: string
          post_id: string
          profile_id: string
          reaction_type: Database["public"]["Enums"]["reaction_type"]
        }
        Insert: {
          created_at?: string
          post_id: string
          profile_id: string
          reaction_type: Database["public"]["Enums"]["reaction_type"]
        }
        Update: {
          created_at?: string
          post_id?: string
          profile_id?: string
          reaction_type?: Database["public"]["Enums"]["reaction_type"]
        }
        Relationships: [
          {
            foreignKeyName: "reactions_post_id_fkey"
            columns: ["post_id"]
            isOneToOne: false
            referencedRelation: "posts"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "reactions_post_id_fkey"
            columns: ["post_id"]
            isOneToOne: false
            referencedRelation: "posts_view"
            referencedColumns: ["post_id"]
          },
          {
            foreignKeyName: "reactions_profile_id_fkey"
            columns: ["profile_id"]
            isOneToOne: false
            referencedRelation: "posts_view"
            referencedColumns: ["user_id"]
          },
          {
            foreignKeyName: "reactions_profile_id_fkey"
            columns: ["profile_id"]
            isOneToOne: false
            referencedRelation: "profiles"
            referencedColumns: ["id"]
          },
        ]
      }
    }
    Views: {
      posts_view: {
        Row: {
          attachment_url: string | null
          avatar_url: string | null
          child_posts: number | null
          content: string | null
          created_at: string | null
          dislikes: number | null
          likes: number | null
          parent_id: string | null
          post_id: string | null
          role: Database["public"]["Enums"]["role_type"] | null
          user_id: string | null
          username: string | null
        }
        Relationships: [
          {
            foreignKeyName: "posts_parent_id_fkey"
            columns: ["parent_id"]
            isOneToOne: false
            referencedRelation: "posts"
            referencedColumns: ["id"]
          },
          {
            foreignKeyName: "posts_parent_id_fkey"
            columns: ["parent_id"]
            isOneToOne: false
            referencedRelation: "posts_view"
            referencedColumns: ["post_id"]
          },
          {
            foreignKeyName: "profiles_id_fkey"
            columns: ["user_id"]
            isOneToOne: true
            referencedRelation: "users"
            referencedColumns: ["id"]
          },
        ]
      }
    }
    Functions: {
      [_ in never]: never
    }
    Enums: {
      reaction_type: "like" | "dislike"
      role_type: "dev" | "verified" | "unverified"
    }
    CompositeTypes: {
      [_ in never]: never
    }
  }
}

type PublicSchema = Database[Extract<keyof Database, "public">]

export type Tables<
  PublicTableNameOrOptions extends
    | keyof (PublicSchema["Tables"] & PublicSchema["Views"])
    | { schema: keyof Database },
  TableName extends PublicTableNameOrOptions extends { schema: keyof Database }
    ? keyof (Database[PublicTableNameOrOptions["schema"]]["Tables"] &
        Database[PublicTableNameOrOptions["schema"]]["Views"])
    : never = never,
> = PublicTableNameOrOptions extends { schema: keyof Database }
  ? (Database[PublicTableNameOrOptions["schema"]]["Tables"] &
      Database[PublicTableNameOrOptions["schema"]]["Views"])[TableName] extends {
      Row: infer R
    }
    ? R
    : never
  : PublicTableNameOrOptions extends keyof (PublicSchema["Tables"] &
        PublicSchema["Views"])
    ? (PublicSchema["Tables"] &
        PublicSchema["Views"])[PublicTableNameOrOptions] extends {
        Row: infer R
      }
      ? R
      : never
    : never

export type TablesInsert<
  PublicTableNameOrOptions extends
    | keyof PublicSchema["Tables"]
    | { schema: keyof Database },
  TableName extends PublicTableNameOrOptions extends { schema: keyof Database }
    ? keyof Database[PublicTableNameOrOptions["schema"]]["Tables"]
    : never = never,
> = PublicTableNameOrOptions extends { schema: keyof Database }
  ? Database[PublicTableNameOrOptions["schema"]]["Tables"][TableName] extends {
      Insert: infer I
    }
    ? I
    : never
  : PublicTableNameOrOptions extends keyof PublicSchema["Tables"]
    ? PublicSchema["Tables"][PublicTableNameOrOptions] extends {
        Insert: infer I
      }
      ? I
      : never
    : never

export type TablesUpdate<
  PublicTableNameOrOptions extends
    | keyof PublicSchema["Tables"]
    | { schema: keyof Database },
  TableName extends PublicTableNameOrOptions extends { schema: keyof Database }
    ? keyof Database[PublicTableNameOrOptions["schema"]]["Tables"]
    : never = never,
> = PublicTableNameOrOptions extends { schema: keyof Database }
  ? Database[PublicTableNameOrOptions["schema"]]["Tables"][TableName] extends {
      Update: infer U
    }
    ? U
    : never
  : PublicTableNameOrOptions extends keyof PublicSchema["Tables"]
    ? PublicSchema["Tables"][PublicTableNameOrOptions] extends {
        Update: infer U
      }
      ? U
      : never
    : never

export type Enums<
  PublicEnumNameOrOptions extends
    | keyof PublicSchema["Enums"]
    | { schema: keyof Database },
  EnumName extends PublicEnumNameOrOptions extends { schema: keyof Database }
    ? keyof Database[PublicEnumNameOrOptions["schema"]]["Enums"]
    : never = never,
> = PublicEnumNameOrOptions extends { schema: keyof Database }
  ? Database[PublicEnumNameOrOptions["schema"]]["Enums"][EnumName]
  : PublicEnumNameOrOptions extends keyof PublicSchema["Enums"]
    ? PublicSchema["Enums"][PublicEnumNameOrOptions]
    : never
