# go-concurrency-key-value-store Lab

### Create a key-value store:
- Store(string key, string value)
- Fetch(string key) -> returns the value
- Start()
- Stop()

### Apply the Actor model:
- As per the slides
- Protected resource: a map holding the data

### Constraints:
- Cannot use sync.Map
- Cannot use sync.Mutex

### Stretch Goals:

#### Could a Command pattern replace the single struct?
- There is a switch-on-type code smell

#### Benchmarking
- Read up on Go benchmark tool
- How does the actor fare?
- Does buffering any channels make a difference?