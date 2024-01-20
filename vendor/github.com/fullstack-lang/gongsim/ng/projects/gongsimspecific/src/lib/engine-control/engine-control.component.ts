import { Component, OnInit, Input } from '@angular/core';

import * as gongsim from 'gongsim'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { Observable, combineLatest, timer } from 'rxjs'
import { GongsimCommandTypeList } from 'gongsim';


@Component({
  selector: 'lib-engine-control',
  templateUrl: './engine-control.component.html',
  styleUrls: ['./engine-control.component.css']
})
export class EngineControlComponent implements OnInit {

  @Input() GONG__StackPath: string = ""

  public engine: gongsim.Engine = new gongsim.Engine
  public engineID: number = 0

  lastEvent: string = ""
  lastEventAgent: string = ""

  nextEventAgent: string = ""
  nextEventName: string = ""
  nextEventTime: string = ""

  engineEventNumber: number = 0

  // animation fof the simulation
  speed = 36;
  clientState = "PAUSED";

  frontRepo = new gongsim.FrontRepo

  // the component is refreshed when modification are performed in the back repo 
  // the checkCommiNbFromBagetCommitNbFromBackTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram
  checkCommiNbFromBagetCommitNbFromBackTimer: Observable<number> = timer(500, 500);
  lastCommitNbFromBack = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  diagramIDForSamocStates: number = 0
  diagramIDForNatoStates: number = 0

  // engineUpdated is the call function
  @Input() engineUpdatedCallbackFunction?: (updateDisplay: boolean) => void
  @Input() simulationTitle: string = ""

  // UpdateDisplay is set to true if the simulation is running or if 
  // the user has provided an input such as AdvanceTillStateChange
  UpdateDisplay = true

  gongsimCommandSingloton: gongsim.GongsimCommand = new gongsim.GongsimCommand

  constructor(
    private frontRepoService: gongsim.FrontRepoService,
    private engineService: gongsim.EngineService,
    private gongsimCommandService: gongsim.GongsimCommandService,
    private commitNbFromBackService: gongsim.CommitNbFromBackService,

    private router: Router) {
  }

  ngOnInit(): void {
    this.commitNbFromBackService.getCommitNbFromBack(500, this.GONG__StackPath).subscribe(
      commiNbFromBagetCommitNbFromBack => {
        if (this.lastCommitNbFromBack < commiNbFromBagetCommitNbFromBack) {

          // console.log("last commit nb " + this.lastCommitNbFromBack + " new: " + commiNbFromBagetCommitNbFromBack)
          this.lastCommitNbFromBack = commiNbFromBagetCommitNbFromBack

          this.frontRepoService.pull(this.GONG__StackPath).subscribe(
            frontRepo => {
              this.frontRepo = frontRepo
              let engines = this.frontRepo.getFrontArray<gongsim.Engine>(gongsim.Engine.GONGSTRUCT_NAME)
              // console.log("Nb of Engines is ", engines.length)

              this.engine = engines[0]

              let commands = this.frontRepo.getFrontArray<gongsim.GongsimCommand>(gongsim.GongsimCommand.GONGSTRUCT_NAME)
              // console.log("Nb of Commands is ", commands.length)

              this.gongsimCommandSingloton = commands[0]

              // this is the callback function from the generic engien to the specific engine 
              if (this.engineUpdatedCallbackFunction) {

                if (this.engine.State == gongsim.EngineState.RUNNING) {
                  this.UpdateDisplay = true
                }
                this.engineUpdatedCallbackFunction(this.UpdateDisplay)

                // reset the need to display updates
                this.UpdateDisplay = false
              }
            }
          )
        }
      }
    )
  }

  fireEventTillStateChange(): void {
    this.gongsimCommandSingloton.Command = gongsim.GongsimCommandType.COMMAND_FIRE_EVENT_TILL_STATES_CHANGE
    this.gongsimCommandSingloton.CommandDate = Date.now().toString()
    this.gongsimCommandService.updateFront(this.gongsimCommandSingloton, this.GONG__StackPath).subscribe(
      gongsimCommand => {
        console.log("FIRE_EVENT_TILL_STATES_CHANGE sent to the backRepo")
      }
    )
    this.UpdateDisplay = true
  }

  fireEvent(): void {
    this.gongsimCommandSingloton.Command = gongsim.GongsimCommandType.COMMAND_FIRE_NEXT_EVENT
    this.gongsimCommandSingloton.CommandDate = Date.now().toString()
    this.gongsimCommandService.updateFront(this.gongsimCommandSingloton, this.GONG__StackPath).subscribe(
      gongsimCommand => {
        console.log("FIRCOMMAND_FIRE_NEXT_EVENT sent to the backRepo")
      }
    )
  }

  reset(): void {
    this.UpdateDisplay = true

    this.gongsimCommandSingloton.Command = gongsim.GongsimCommandType.COMMAND_RESET
    this.gongsimCommandSingloton.CommandDate = Date.now().toString()
    this.gongsimCommandService.updateFront(this.gongsimCommandSingloton, this.GONG__StackPath).subscribe(
      gongsimCommand => {
        console.log("RESCOMMAND_RESET sent to the backRepo")
      }
    )
  }

  play(): void {
    this.UpdateDisplay = true

    this.gongsimCommandSingloton.Command = gongsim.GongsimCommandType.COMMAND_PLAY
    this.gongsimCommandSingloton.CommandDate = Date.now().toString()
    this.gongsimCommandService.updateFront(this.gongsimCommandSingloton, this.GONG__StackPath).subscribe(
      gongsimCommand => {
        console.log("PLAY sent to the backRepo")
      }
    )
  }

  pause(): void {
    this.gongsimCommandSingloton.Command = gongsim.GongsimCommandType.COMMAND_PAUSE
    this.gongsimCommandSingloton.CommandDate = Date.now().toString()
    this.gongsimCommandService.updateFront(this.gongsimCommandSingloton, this.GONG__StackPath).subscribe(
      gongsimCommand => {
        console.log("PAUSE sent to the backRepo")
      }
    )
  }

  increaseSpeed100percent(): void {
    this.gongsimCommandSingloton.SpeedCommandType = gongsim.SpeedCommandType.COMMAND_INCREASE_SPEED_100_PERCENTS
    this.gongsimCommandSingloton.DateSpeedCommand = Date.now().toString()
    this.gongsimCommandService.updateFront(this.gongsimCommandSingloton, this.GONG__StackPath).subscribe(
      gongsimCommand => {
        console.log("INCCOMMAND_INCREASE_SPEED_100_PERCENTS sent to the backRepo")
      }
    )

    this.UpdateDisplay = true
  }

  decreaseSpeed50percent(): void {
    this.gongsimCommandSingloton.SpeedCommandType = gongsim.SpeedCommandType.COMMAND_DECREASE_SPEED_50_PERCENTS
    this.gongsimCommandSingloton.DateSpeedCommand = Date.now().toString()
    this.gongsimCommandService.updateFront(this.gongsimCommandSingloton, this.GONG__StackPath).subscribe(
      gongsimCommand => {
        console.log("DECCOMMAND_DECREASE_SPEED_50_PERCENTS sent to the backRepo")
      }
    )
    this.UpdateDisplay = true
  }

  publishAgentStates(): void {
    this.UpdateDisplay = true
  }

  publishAdvance10Minutes(): void {
    this.gongsimCommandSingloton.Command = gongsim.GongsimCommandType.COMMAND_ADVANCE_10_MIN
    this.gongsimCommandSingloton.CommandDate = Date.now().toString()
    this.gongsimCommandService.updateFront(this.gongsimCommandSingloton, this.GONG__StackPath).subscribe(
      gongsimCommand => {
        console.log("ADCOMMAND_ADVANCE_10_MIN sent to the backRepo")
      }
    )
    this.UpdateDisplay = true
  }
}
