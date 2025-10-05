INSERT INTO decks (owner_id, name) VALUES
  ((SELECT id FROM users WHERE username = 'pusiapaker'), 'Russian 1 - Golosa Chapter 1'),
  ((SELECT id FROM users WHERE username = 'pusiapaker'), 'Russian 1 - Golosa Chapter 2'),
  ((SELECT id FROM users WHERE username = 'morgan'), 'Aligator Facts'),
  ((SELECT id FROM users WHERE username = 'allie'), 'Cool Plants and shitz'),
  ((SELECT id FROM users WHERE username = 'allie'), 'Variable Length Cards')
ON CONFLICT DO NOTHING;


-- Russian 1 - Golosa Chapter 1
INSERT INTO cards (deck_id, term, definition) VALUES
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 1'), 'Привет', 'Hello'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 1'), 'Спасибо', 'Thank you'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 1'), 'Пожалуйста', 'Please / You’re welcome'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 1'), 'Да', 'Yes'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 1'), 'Нет', 'No'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 1'), 'Извините', 'Excuse me / Sorry'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 1'), 'Как дела?', 'How are you?'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 1'), 'Хорошо', 'Good / Fine'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 1'), 'Плохо', 'Bad'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 1'), 'До свидания', 'Goodbye')
ON CONFLICT DO NOTHING;

-- Russian 1 - Golosa Chapter 2
INSERT INTO cards (deck_id, term, definition) VALUES
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 2'), 'Мама', 'Mother'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 2'), 'Папа', 'Father'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 2'), 'Брат', 'Brother'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 2'), 'Сестра', 'Sister'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 2'), 'Дом', 'House'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 2'), 'Квартира', 'Apartment'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 2'), 'Улица', 'Street'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 2'), 'Школа', 'School'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 2'), 'Студент', 'Student (male)'),
((SELECT id FROM decks WHERE name = 'Russian 1 - Golosa Chapter 2'), 'Студентка', 'Student (female)')
ON CONFLICT DO NOTHING;

-- Aligator Facts
INSERT INTO cards (deck_id, term, definition) VALUES
((SELECT id FROM decks WHERE name = 'Aligator Facts'), 'Lifespan', 'American alligators can live 35-50 years in the wild'),
((SELECT id FROM decks WHERE name = 'Aligator Facts'), 'Size', 'Males can grow over 13 feet (4 m) long'),
((SELECT id FROM decks WHERE name = 'Aligator Facts'), 'Diet', 'They are carnivorous, eating fish, birds, mammals, and reptiles'),
((SELECT id FROM decks WHERE name = 'Aligator Facts'), 'Habitat', 'Alligators prefer freshwater swamps, marshes, and rivers'),
((SELECT id FROM decks WHERE name = 'Aligator Facts'), 'Teeth', 'Alligators have about 80 teeth, which are replaced throughout life'),
((SELECT id FROM decks WHERE name = 'Aligator Facts'), 'Speed', 'They can run up to 11 mph (18 km/h) on land in short bursts'),
((SELECT id FROM decks WHERE name = 'Aligator Facts'), 'Bite Force', 'Measured at over 2,000 pounds per square inch'),
((SELECT id FROM decks WHERE name = 'Aligator Facts'), 'Eggs', 'Females lay 20-50 eggs at a time in a nest of vegetation'),
((SELECT id FROM decks WHERE name = 'Aligator Facts'), 'Sex Determination', 'Egg temperature decides sex: warmer makes males'),
((SELECT id FROM decks WHERE name = 'Aligator Facts'), 'Conservation', 'They were once endangered but recovered through protections')
ON CONFLICT DO NOTHING;

-- Cool Plants and shitz
INSERT INTO cards (deck_id, term, definition) VALUES
((SELECT id FROM decks WHERE name = 'Cool Plants and shitz'), 'Venus Flytrap', 'Carnivorous plant that snaps shut on insects'),
((SELECT id FROM decks WHERE name = 'Cool Plants and shitz'), 'Photosynthesis', 'Process plants use to make energy from sunlight'),
((SELECT id FROM decks WHERE name = 'Cool Plants and shitz'), 'Bamboo', 'Fastest-growing plant, can grow up to 3 feet per day'),
((SELECT id FROM decks WHERE name = 'Cool Plants and shitz'), 'Succulent', 'Water-storing plant adapted to arid climates'),
((SELECT id FROM decks WHERE name = 'Cool Plants and shitz'), 'Ferns', 'Ancient plants that reproduce via spores'),
((SELECT id FROM decks WHERE name = 'Cool Plants and shitz'), 'Baobab Tree', 'Tree with a trunk that stores thousands of liters of water'),
((SELECT id FROM decks WHERE name = 'Cool Plants and shitz'), 'Aloe Vera', 'Succulent known for soothing burns'),
((SELECT id FROM decks WHERE name = 'Cool Plants and shitz'), 'Corpse Flower', 'Rare plant with a strong odor of rotting flesh'),
((SELECT id FROM decks WHERE name = 'Cool Plants and shitz'), 'Moss', 'Non-vascular plant that absorbs water directly'),
((SELECT id FROM decks WHERE name = 'Cool Plants and shitz'), 'Orchid', 'Large family of flowering plants with diverse shapes and colors')
ON CONFLICT DO NOTHING;

INSERT INTO cards (deck_id, term, definition) VALUES
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'The quick brown fox jumps over several extraordinarily lazy neighborhood dogs during sunrise',
 'A whimsical sentence constructed to exercise typography and layout, this phrase contains varied word lengths and spacing. It ensures our rendering pipeline handles kerning, wrapping, and overflow without clipping, jitter, or misalignment when presented across responsive breakpoints.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Considerations for resilient database connection pooling under unpredictable bursty traffic conditions',
 'Connection pools must guard against thundering herds, slow query amplification, and transient network faults. Employ jittered backoff, health checks, circuit breakers, and sane pool limits to avoid saturation. Observability with per-query timing and pool metrics is essential for timely remediation.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Semantic versioning nuances when libraries expose unstable experimental interfaces to early adopters',
 'When APIs remain experimental, communicate this status clearly and avoid implying stability via version numbers alone. Use feature flags, pre-release identifiers, and thorough changelogs. Breaking changes should be deliberate, documented, and accompanied by migration notes and deprecation timelines.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Techniques for minimizing cumulative layout shift in content heavy client rendered applications',
 'Prevent unexpected movement by reserving space for media, using intrinsic size attributes, avoiding ad reflow, and deferring noncritical fonts. Measure layout shift metrics continuously and set budgets. Stabilize skeletons and loaders so users do not lose reading position or tap incorrect controls.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Designing idempotent request handlers that remain correct across retries, timeouts, and duplicates',
 'Idempotence prevents duplicate side effects when clients retry. Use unique request keys, conditional updates, and transactional semantics. Store processed tokens, prefer PUT for upserts, and structure operations so repeated execution yields the same result. Log correlation IDs for traceability.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Human readable yet globally unique identifiers suited for URLs, logs, and database keys at scale',
 'Choose identifiers that balance collision resistance, sortability, and ergonomics. ULIDs and UUIDv7 provide time-ordered properties useful for pagination and storage locality. Avoid leaking entropy sources. Document formats and validate strictly to keep systems debuggable and interoperable.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Strategies for progressive enhancement when JavaScript fails, stalls, or is deliberately disabled',
 'Deliver meaningful HTML first, then hydrate capabilities as resources load. Provide accessible navigation, forms that submit server-side, and content that remains readable without scripts. Enhance, do not replace. Fail gracefully with clear messages and minimal dependency on fragile client states.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Effective log message style guides that reduce noise while increasing incident investigation speed',
 'Logs should be structured, leveled, and contextual. Include stable keys, correlation IDs, and precise causes rather than vague statements. Avoid personal data and secrets. Prefer one event per line with machine parsable fields. Consistency across services accelerates queries and root cause analysis.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Caching pitfalls involving stale writes, phantom reads, and incoherent replicas across regions',
 'Caches accelerate reads yet complicate correctness. Apply explicit TTLs, cache busting on writes, and validation headers. For multi-region systems, prefer write-through strategies or versioned keys. Monitor hit ratios, eviction causes, and tail latencies to detect subtle coherence regressions early.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Unicode normalization and grapheme cluster boundaries for truly user friendly text handling',
 'Text that appears as one character may consist of several code points. Normalize inputs, iterate by grapheme clusters, and respect locale rules for case and collation. Mishandling diacritics, emoji skin tones, and zero width joiners can break search, truncation, or cursor navigation unexpectedly.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Threat modeling focused on misuse cases, least privilege boundaries, and safe secret management',
 'Identify assets, entry points, and attacker goals. Enforce least privilege at every hop, rotate credentials, and store secrets in dedicated vaults. Assume compromise and design blast radius limits. Document trust boundaries explicitly so reviewers can reason about authentication and authorization flows.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Latency budgets for end to end interactions spanning mobile radio, CDN edges, and origin services',
 'Define a strict budget that allocates milliseconds to DNS, TLS, TTFB, and rendering. Measure real devices on real networks, not laboratory ideals. Optimize critical path resources, compress thoughtfully, and push personalization off the hot path. Report p50, p90, and p99 to capture tail behaviors.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Database migration playbook covering forward only changes and reversible safety valves',
 'Prefer additive migrations, deploy code that tolerates both schemas, and only then remove deprecated fields. Guard destructive steps behind feature flags and run online operations to avoid long locks. Keep rollback levers, backups, and rehearse the plan in staging with production-like datasets.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Approaches for testing accessibility including keyboard flows and assistive technology audits',
 'Automate linting for ARIA usage, color contrast, and heading structure. Manually test keyboard traps, focus order, and skip links. Run screen readers to verify labels, roles, and live regions. Accessibility is a cross-cutting quality attribute, not a checkbox added at the end of development.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Operational excellence through runbooks, on call rotations, and blameless post incident reviews',
 'Clear runbooks reduce cognitive load during stress. Rotations must be humane with sustainable alert policies. After incidents, analyze contributing factors without blame, fix systemic issues, and document learnings. Track action items to completion so the same failure mode does not reappear quietly.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'API pagination patterns that scale while avoiding duplicate, missing, or reordered results',
 'Offset based pagination is simple but unstable under concurrent writes. Prefer cursor based approaches using monotonic sort keys. Return next links and stable page sizes. Document consistency guarantees. Clients should tolerate gaps and deduplicate by IDs to prevent subtle reconciliation errors.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Choosing effective metrics, SLIs, SLOs, and error budgets to guide pragmatic reliability work',
 'Select user centric indicators such as availability, latency, and correctness. Define SLOs with realistic targets and enforce them through error budgets. When budgets burn, pause risky launches and prioritize stabilization. Visualize long term trends so reliability investment remains data driven.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Internationalization concerns beyond translation, including pluralization, calendars, and numerals',
 'Support locale specific plural rules, date systems, and numbering traditions. Avoid hard coded assumptions about week starts, name order, or currency formats. Store canonical data and render per locale at the edge. Provide translators with context so phrasing remains accurate and culturally respectful.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Data retention policies that balance analytics value with privacy, compliance, and storage cost',
 'Collect only what you can protect and justify. Define lifetimes for raw and derived data, then enforce deletion automatically. Anonymize when possible and document lawful bases. Regularly review purpose creep and access scopes. Storage is cheap until subpoenas, breaches, or reputational harm arrive.'
),
((SELECT id FROM decks WHERE name = 'Variable Length Cards'),
 'Holistic performance profiling that spans CPU, memory, I O, and garbage collection interactions',
 'Performance issues rarely live in one layer. Profile hot paths with realistic workloads, capture flame graphs, and inspect allocation rates. Watch garbage collector pauses, disk contention, and lock contention together. Optimization should follow evidence, not folklore, and must include regression tests.'
)
ON CONFLICT DO NOTHING;