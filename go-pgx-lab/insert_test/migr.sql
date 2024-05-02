
CREATE TABLE trial.items_name_btree (
    id uuid NOT NULL,
    name text NOT NULL DEFAULT ''::text,
    number numeric,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz,
    standard_id uuid NOT NULL DEFAULT '63b14b45-19d5-49d7-99dd-2a09e1b2f82f'::uuid,
    PRIMARY KEY (id)
);

CREATE INDEX idx_trial_items_name_btree ON trial.items_name_btree(name);

---

CREATE TABLE trial.items_name_gin (
    id uuid NOT NULL,
    name text NOT NULL DEFAULT ''::text,
    number numeric,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz,
    standard_id uuid NOT NULL DEFAULT '63b14b45-19d5-49d7-99dd-2a09e1b2f82f'::uuid,
    PRIMARY KEY (id)
);

CREATE INDEX idx_trial_items_name_gin ON trial.items_name_gin USING gin (name gin_trgm_ops);

--- 

CREATE INDEX idx_trial_items_name_btree_gin ON trial.items_name_btree_gin USING gin (name gin_trgm_ops);