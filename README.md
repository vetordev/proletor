# proletor
Bot manager running jobs

                         +------------------+
                         |                  |
                         |                  |
              +----------+       Main       +-----------+
              |          |                  |           |
              |          |                  |           |
              |          +------------------+           |
              |Start                                    |Start
              |                                         |
              |                                         |
    +---------v---------+                     +---------v--------+
    |                   |                     |                  |
    |                   |                     |                  |
    |      Service      |                     |      Server      |
    |                   |                     |                  |
    |                   |                     |                  |
    +-------------------+                     +------------------+
      Load bots and run                         Provide a api to
                                                setup new bots
    