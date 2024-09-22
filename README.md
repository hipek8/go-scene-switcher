# Scene switcher

Features:
* Automatically sync TV to Amplituner status
   - When Amplituner turns off, the TV turns off
   - When Amplituner is on and scene set to PC the TV turns on
* Set scene via Rest API `/scene/<scene>`
  - off
  - pc
  - rns
  - r357
  - tv
  - interface
  - mixer
* Handle specific devices `/device/<device>/<command>`
  - fan
    * toggle
    * oscillate 
    * \+
    * \-
  - tv
    * on
    * off
  - mc
    * on
    - off
    - pc
    - net_radio
    - rns
    - r357
    - vol/low
    - vol/mid
    - vol/high
    - vol/xxl