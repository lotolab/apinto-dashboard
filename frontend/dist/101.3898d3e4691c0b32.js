"use strict";(self.webpackChunkapinto=self.webpackChunkapinto||[]).push([[101],{9538:(N,z,l)=>{l.d(z,{m:()=>p});var A=l(4006);function p(a,g){Object.keys(a.controls).forEach(m=>{a.controls[m]instanceof A.cw?p(a.controls[m],g[m]):a.controls[m].setValue(g[m])})}},7101:(N,z,l)=>{l.r(z),l.d(z,{SystemModule:()=>ue});var A=l(6895),p=l(9132),a=l(4006),g=l(9538),m=l(1461);const D=[{label:"POST",value:"POST"},{label:"GET",value:"GET"}],L=[{label:"JSON",value:"JSON"},{label:"form-data",value:"form-data"}],U=[{label:"\u5355\u6b21\u53d1\u9001",value:"single"},{label:"\u591a\u6b21\u53d1\u9001",value:"many"}],I=[{label:"\u4e0d\u8bbe\u7f6e\u4efb\u4f55\u534f\u8bae",value:"none"},{label:"SSL\u534f\u8bae",value:"ssl"},{label:"TLS\u534f\u8bae",value:"tls"}],J=[{title:"\u5e94\u7528\u540d\u79f0"},{title:"\u5e94\u7528ID"},{title:"\u9274\u6743Token"},{title:"\u5173\u8054\u6807\u7b7e"},{title:"\u7981\u7528\u72b6\u6001",width:84},{title:"\u66f4\u65b0\u8005"},{title:"\u66f4\u65b0\u65f6\u95f4",showSort:!0},{title:"\u64cd\u4f5c",right:!0}],M=[{key:"name",copy:!0},{key:"id",copy:!0},{key:"token",copy:!0},{key:"tags"},{key:"status"},{key:"operator"},{key:"updateTime"},{type:"btn",right:!0,btns:[{title:"\u66f4\u65b0\u9274\u6743Token"},{title:"\u590d\u5236Token"},{title:"\u67e5\u770b"},{title:"\u5220\u9664"}]}];var e=l(4650),k=l(2266),h=l(3),C=l(4963),F=l(2366),f=l(4162),b=l(3679),u=l(6704),y=l(999),v=l(8344),c=l(9561);let O=(()=>{class n{constructor(t,o,s,r){this.fb=t,this.appConfigService=o,this.api=s,this.message=r,this.editPage=!1,this.validateForm=new a.cw({}),this.nzDisabled=!1,this.emailId="",this.autoTips=m.Sm,this.listOfProtocols=[...I],this.validateForm=this.fb.group({smtpUrl:["",[a.kI.required]],smtpPort:[null,[a.kI.required]],protocol:["ssl",[a.kI.required]],email:["",[a.kI.email]],account:[""],password:[""]}),this.appConfigService.reqFlashBreadcrumb([{title:"\u90ae\u7bb1\u8bbe\u7f6e"}])}ngOnInit(){this.getEmailMessage()}disabledEdit(t){this.nzDisabled=t}getEmailMessage(){this.api.get("email").subscribe(t=>{0===t.code&&t.data.emailInfo&&(t.data.emailInfo.protocol=""===t.data.emailInfo.protocol?"none":t.data.emailInfo.protocol,(0,g.m)(this.validateForm,t.data.emailInfo),this.emailId=t.data.emailInfo.uuid,this.editPage=!0)})}save(){if(this.validateForm.valid){const t={...this.validateForm.value};"none"===t.protocol&&(t.protocol=""),this.editPage?this.api.put("email",{...t,uuid:this.emailId}).subscribe(o=>{0===o.code&&this.message.success(o.msg||"\u4fee\u6539\u901a\u77e5\u90ae\u7bb1\u6210\u529f\uff01")}):this.api.post("email",{...t}).subscribe(o=>{0===o.code&&this.message.success(o.msg||"\u521b\u5efa\u901a\u77e5\u90ae\u7bb1\u6210\u529f\uff01")})}else Object.values(this.validateForm.controls).forEach(t=>{t.invalid&&(t.markAsDirty(),t.updateValueAndValidity({onlySelf:!0}))})}}return n.\u0275fac=function(t){return new(t||n)(e.Y36(a.QS),e.Y36(k.a),e.Y36(h.s),e.Y36(C.A))},n.\u0275cmp=e.Xpm({type:n,selectors:[["eo-ng-system-email-config"]],decls:38,vars:22,consts:[[1,"m-btnbase","mb-0"],["nzType","info","nzMessage","\u90ae\u7bb1\u8bbe\u7f6e\u7528\u4e8e\u89e6\u53d1\u544a\u8b66\u7b56\u7565\u7cfb\u7edf\u53d1\u9001\u90ae\u4ef6\u7ed9\u7528\u6237\u3002","nzShowIcon","",3,"nzBanner","nzCloseable"],[1,"mt-mbase"],["nz-form","","autocomplete","off",3,"nzNoColon","nzAutoTips","formGroup"],["nzRequired","",3,"nzSpan"],[3,"nzSpan"],["type","text","eo-ng-input","","name","smtpUrl","formControlName","smtpUrl","placeholder","\u8bf7\u8f93\u5165","required","","eoNgUserAccess","system/email",1,"w-INPUT_NORMAL",3,"disabledEdit"],["type","number","eo-ng-input","","name","smtpPort","formControlName","smtpPort","placeholder","\u8bf7\u8f93\u5165","required","","eoNgUserAccess","system/email",1,"w-INPUT_NORMAL"],["formControlName","protocol","nzPlaceHolder","\u8bf7\u9009\u62e9",1,"w-INPUT_NORMAL",3,"nzOptions","nzDisabled"],["type","text","eo-ng-input","","name","email","formControlName","email","placeholder","\u8bf7\u8f93\u5165","eoNgUserAccess","system/email",1,"w-INPUT_NORMAL"],["eo-ng-input","","name","account","placeholder","\u8bf7\u8f93\u5165","formControlName","account","eoNgUserAccess","system/email",1,"w-INPUT_NORMAL"],[1,"mb-0"],["eo-ng-input","","name","password","placeholder","\u8bf7\u8f93\u5165","formControlName","password","eoNgUserAccess","system/email",1,"w-INPUT_NORMAL"],[1,"bg-white","mb-0","sticky","bottom-0","pt-mbase","pb-mbase"],[3,"nzOffset","nzSpan"],["eoNgUserAccess","system/email","type","submit","eo-ng-button","",1,"ant-btn-primary",3,"click"]],template:function(t,o){1&t&&(e.TgZ(0,"div",0),e._UZ(1,"eo-ng-feedback-alert",1),e.qZA(),e.TgZ(2,"div",2)(3,"form",3)(4,"nz-form-item")(5,"nz-form-label",4),e._uU(6,"SMTP\u5730\u5740\uff1a"),e.qZA(),e.TgZ(7,"nz-form-control",5)(8,"input",6),e.NdJ("disabledEdit",function(r){return o.disabledEdit(r)}),e.qZA()()(),e.TgZ(9,"nz-form-item")(10,"nz-form-label",4),e._uU(11,"SMTP\u7aef\u53e3\uff1a"),e.qZA(),e.TgZ(12,"nz-form-control",5),e._UZ(13,"input",7),e.qZA()(),e.TgZ(14,"nz-form-item")(15,"nz-form-label",4),e._uU(16,"\u901a\u4fe1\u534f\u8bae\uff1a"),e.qZA(),e.TgZ(17,"nz-form-control",5),e._UZ(18,"eo-ng-select",8),e.qZA()(),e.TgZ(19,"nz-form-item")(20,"nz-form-label",5),e._uU(21,"\u53d1\u4ef6\u90ae\u7bb1\uff1a"),e.qZA(),e.TgZ(22,"nz-form-control",5),e._UZ(23,"input",9),e.qZA()(),e.TgZ(24,"nz-form-item")(25,"nz-form-label",5),e._uU(26,"\u8d26\u53f7\uff1a"),e.qZA(),e.TgZ(27,"nz-form-control",5),e._UZ(28,"input",10),e.qZA()(),e.TgZ(29,"nz-form-item",11)(30,"nz-form-label",5),e._uU(31,"\u5bc6\u7801\uff1a"),e.qZA(),e.TgZ(32,"nz-form-control",5),e._UZ(33,"input",12),e.qZA()(),e.TgZ(34,"nz-form-item",13)(35,"nz-form-control",14)(36,"button",15),e.NdJ("click",function(){return o.save()}),e._uU(37),e.qZA()()()()()),2&t&&(e.xp6(1),e.Q6J("nzBanner",!0)("nzCloseable",!0),e.xp6(2),e.Q6J("nzNoColon",!0)("nzAutoTips",o.autoTips)("formGroup",o.validateForm),e.xp6(2),e.Q6J("nzSpan",9),e.xp6(2),e.Q6J("nzSpan",10),e.xp6(3),e.Q6J("nzSpan",9),e.xp6(2),e.Q6J("nzSpan",10),e.xp6(3),e.Q6J("nzSpan",9),e.xp6(2),e.Q6J("nzSpan",10),e.xp6(1),e.Q6J("nzOptions",o.listOfProtocols)("nzDisabled",o.nzDisabled),e.xp6(2),e.Q6J("nzSpan",9),e.xp6(2),e.Q6J("nzSpan",10),e.xp6(3),e.Q6J("nzSpan",9),e.xp6(2),e.Q6J("nzSpan",10),e.xp6(3),e.Q6J("nzSpan",9),e.xp6(2),e.Q6J("nzSpan",10),e.xp6(3),e.Q6J("nzOffset",9)("nzSpan",18),e.xp6(2),e.hij(" ",o.editPage?"\u63d0\u4ea4":"\u4fdd\u5b58"," "))},dependencies:[a._Y,a.Fj,a.wV,a.JJ,a.JL,a.Q7,a.sg,a.u,F.oO,f.qK,b.t3,b.SK,u.Lr,u.Nx,u.iK,u.Fd,y.g,v.Hf,c.$Q],encapsulation:2}),n})(),Q=(()=>{class n{}return n.\u0275fac=function(t){return new(t||n)},n.\u0275cmp=e.Xpm({type:n,selectors:[["eo-ng-system-email"]],decls:1,vars:0,template:function(t,o){1&t&&e._UZ(0,"router-outlet")},dependencies:[p.lC],encapsulation:2}),n})(),_=(()=>{class n{constructor(t,o,s,r,d){this.message=t,this.api=o,this.router=s,this.fb=r,this.appConfigService=d,this.editPage=!1,this.appId="",this.validateForm=new a.cw({}),this.submitButtonLoading=!1,this.autoTips=m.Sm,this.nzDisabled=!1,this.appConfigService.reqFlashBreadcrumb([{title:"\u5916\u90e8\u5e94\u7528"},{title:"\u65b0\u5efa\u5916\u90e8\u5e94\u7528"}]),this.validateForm=this.fb.group({name:["",[a.kI.required]],id:[""],desc:[""]})}ngOnInit(){this.editPage?(this.getApplicationMessage(),this.appConfigService.reqFlashBreadcrumb([{title:"\u5916\u90e8\u5e94\u7528"},{title:"\u5916\u90e8\u5e94\u7528\u8be6\u60c5"}])):this.getApplicationId()}getApplicationMessage(){this.api.get("external-app",{id:this.appId}).subscribe(t=>{0===t.code&&(this.validateForm.controls.name.setValue(t.data.info.name),this.validateForm.controls.id.setValue(t.data.info.id),this.validateForm.controls.desc.setValue(t.data.info.desc),this.validateForm.controls.id.disable())})}getApplicationId(){this.api.get("random/external-app/id").subscribe(t=>{0===t.code&&this.validateForm.controls.id.setValue(t.data.id)})}disabledEdit(t){this.nzDisabled=t}saveApplication(){this.validateForm.valid?(this.submitButtonLoading=!0,this.editPage?this.api.put("external-app",{...this.validateForm.value},{id:this.validateForm.controls.id.value}).subscribe(t=>{this.submitButtonLoading=!1,0===t.code&&(this.message.success(t.msg||"\u4fee\u6539\u6210\u529f"),this.backToList())}):this.api.post("external-app",{...this.validateForm.value}).subscribe(t=>{this.submitButtonLoading=!1,0===t.code&&(this.message.success(t.msg||"\u6dfb\u52a0\u6210\u529f"),this.backToList())})):Object.values(this.validateForm.controls).forEach(t=>{t.invalid&&(t.markAsDirty(),t.updateValueAndValidity({onlySelf:!0}))})}backToList(){this.router.navigate(["/","system","ext-app"])}}return n.\u0275fac=function(t){return new(t||n)(e.Y36(c.VY),e.Y36(h.s),e.Y36(p.F0),e.Y36(a.QS),e.Y36(k.a))},n.\u0275cmp=e.Xpm({type:n,selectors:[["eo-ng-external-app-create"]],inputs:{editPage:"editPage",appId:"appId"},decls:23,vars:13,consts:[[1,"mt-formtop"],["nz-form","","autocomplete","off",3,"nzNoColon","nzAutoTips","formGroup"],["nzRequired","",3,"nzSpan"],[3,"nzSpan"],["eo-ng-input","","required","","name","name","placeholder","\u8bf7\u8f93\u5165","formControlName","name","eoNgUserAccess","system/ext-app",1,"w-INPUT_NORMAL",3,"disabledEdit"],["eo-ng-input","","required","","name","id","placeholder","\u8bf7\u8f93\u5165","formControlName","id","eoNgUserAccess","system/ext-app",1,"w-INPUT_NORMAL"],[1,"mb-0"],["name","desc","rows","3","eo-ng-input","","formControlName","desc","placeholder","\u8bf7\u8f93\u5165","eoNgUserAccess","system/ext-app",1,"w-INPUT_NORMAL"],[1,"mb-0","sticky","bg-white","bottom-0","py-mbase","z-50"],[3,"nzOffset","nzSpan"],["type","submit","eo-ng-button","","eoNgUserAccess","system/ext-app",1,"ant-btn-primary",3,"nzLoading","click"],["eo-ng-button","",1,"ml-btnybase",3,"click"]],template:function(t,o){1&t&&(e.TgZ(0,"div",0)(1,"form",1)(2,"nz-form-item")(3,"nz-form-label",2),e._uU(4,"\u5e94\u7528\u540d\u79f0\uff1a"),e.qZA(),e.TgZ(5,"nz-form-control",3)(6,"input",4),e.NdJ("disabledEdit",function(r){return o.disabledEdit(r)}),e.qZA()()(),e.TgZ(7,"nz-form-item")(8,"nz-form-label",2),e._uU(9,"\u5e94\u7528ID\uff1a"),e.qZA(),e.TgZ(10,"nz-form-control",3),e._UZ(11,"input",5),e.qZA()(),e.TgZ(12,"nz-form-item",6)(13,"nz-form-label",3),e._uU(14,"\u63cf\u8ff0\uff1a"),e.qZA(),e.TgZ(15,"nz-form-control",3),e._UZ(16,"textarea",7),e.qZA()(),e.TgZ(17,"nz-form-item",8)(18,"nz-form-control",9)(19,"button",10),e.NdJ("click",function(){return o.saveApplication()}),e._uU(20),e.qZA(),e.TgZ(21,"button",11),e.NdJ("click",function(){return o.backToList()}),e._uU(22," \u53d6\u6d88 "),e.qZA()()()()()),2&t&&(e.xp6(1),e.Q6J("nzNoColon",!0)("nzAutoTips",o.autoTips)("formGroup",o.validateForm),e.xp6(2),e.Q6J("nzSpan",7),e.xp6(2),e.Q6J("nzSpan",12),e.xp6(3),e.Q6J("nzSpan",7),e.xp6(2),e.Q6J("nzSpan",12),e.xp6(3),e.Q6J("nzSpan",7),e.xp6(2),e.Q6J("nzSpan",12),e.xp6(3),e.Q6J("nzOffset",7)("nzSpan",12),e.xp6(1),e.Q6J("nzLoading",o.submitButtonLoading),e.xp6(1),e.hij(" ",o.editPage?"\u63d0\u4ea4":"\u4fdd\u5b58"," "))},dependencies:[a._Y,a.Fj,a.JJ,a.JL,a.Q7,a.sg,a.u,F.oO,f.qK,b.t3,b.SK,u.Lr,u.Nx,u.iK,u.Fd,y.g],encapsulation:2}),n})(),P=(()=>{class n{}return n.\u0275fac=function(t){return new(t||n)},n.\u0275cmp=e.Xpm({type:n,selectors:[["eo-ng-external-app"]],decls:1,vars:0,template:function(t,o){1&t&&e._UZ(0,"router-outlet")},dependencies:[p.lC],encapsulation:2}),n})();var S=l(7504),E=l(2663),x=l(8002),Z=l(9041);const q=["switchTpl"],W=["copyTokenTpl"];function H(n,i){if(1&n){const t=e.EpF();e.TgZ(0,"button",7),e.NdJ("copyCallback",function(){e.CHM(t);const s=e.oxw();return e.KtG(s.callback())})("click",function(s){return s.stopPropagation()}),e.TgZ(1,"span",8),e.O4$(),e.TgZ(2,"svg",9),e._UZ(3,"use",10),e.qZA()()()}2&n&&(e.Q6J("copyText",i.item.data.token),e.xp6(1),e.Q6J("nzTooltipVisible",!1))}function R(n,i){if(1&n){const t=e.EpF();e.TgZ(0,"eo-ng-switch",11),e.NdJ("ngModelChange",function(s){const d=e.CHM(t).item;return e.KtG(d.statusBoolean=s)})("click",function(s){const d=e.CHM(t).item,T=e.oxw();return e.KtG(T.disabledApp(d,s))}),e.qZA()}if(2&n){const t=i.item,o=e.oxw();e.Q6J("ngModel",t.statusBoolean)("nzControl",!0)("nzDisabled",o.nzDisabled)}}let Y=(()=>{class n{constructor(t,o,s,r,d){this.router=t,this.message=o,this.modalService=s,this.api=r,this.appConfigService=d,this.nzDisabled=!1,this.appsTableHeadName=[...J],this.appsTableBody=[...M],this.appsList=[],this.appsTableClick=T=>{this.getAppMessage(T.data)},this.deleteApp=T=>{this.api.delete("external-app",{id:T.id}).subscribe(w=>{0===w.code&&(this.message.success(w.msg||"\u5220\u9664\u6210\u529f!"),this.getAppsData())})},this.callback=()=>{this.message.success("\u590d\u5236\u6210\u529f")},this.appConfigService.reqFlashBreadcrumb([{title:"\u5916\u90e8\u5e94\u7528"}])}ngOnInit(){this.getAppsData()}ngAfterViewInit(){this.appsTableBody[4].title=this.switchTpl,this.appsTableBody[7].btns[0].disabledFn=()=>this.nzDisabled,this.appsTableBody[7].btns[0].click=t=>{this.updateToken(t.data)},this.appsTableBody[7].btns[1].type=this.copyTokenTpl,this.appsTableBody[7].btns[2].click=t=>{this.getAppMessage(t.data)},this.appsTableBody[7].btns[3].disabledFn=()=>this.nzDisabled,this.appsTableBody[7].btns[3].click=t=>{this.deleteAppModal(t.data)}}disabledEdit(t){this.nzDisabled=t}getAppsData(){this.api.get("external-apps").subscribe(t=>{if(0===t.code){this.appsList=t.data.apps;for(const o in this.appsList)this.appsList[o].statusBoolean=2===this.appsList[o].status}})}getAppMessage(t){this.router.navigate(["/","system","ext-app","message",t.id])}addApp(){this.router.navigate(["/","system","ext-app","create"])}updateToken(t){this.api.put("external-app/token",null,{id:t.id}).subscribe(o=>{0===o.code&&(this.message.success(o.msg||"\u5237\u65b0\u9274\u6743Token\u6210\u529f\uff01"),this.getAppsData())})}disabledApp(t,o){o.stopPropagation();let s="";s=t.statusBoolean?"external-app/enable":"external-app/disable",this.api.put(s,null,{id:t.id}).subscribe(r=>{0===r.code&&(t.statusBoolean=!t.statusBoolean,this.message.success(r.msg||("disable"===s.split("/")[1]?"\u7981\u7528":"\u542f\u7528")+"\u6210\u529f\uff01"),this.getAppsData())})}deleteAppModal(t){this.modalService.create({nzTitle:"\u5220\u9664",nzContent:"\u8be5\u6570\u636e\u5220\u9664\u540e\u5c06\u65e0\u6cd5\u627e\u56de\uff0c\u8bf7\u786e\u8ba4\u662f\u5426\u5220\u9664\uff1f",nzClosable:!0,nzClassName:"delete-modal",nzWidth:S.G_,nzOkDanger:!0,nzOnOk:()=>{this.deleteApp(t)}})}}return n.\u0275fac=function(t){return new(t||n)(e.Y36(p.F0),e.Y36(c.VY),e.Y36(c._B),e.Y36(h.s),e.Y36(k.a))},n.\u0275cmp=e.Xpm({type:n,selectors:[["eo-ng-external-app-list"]],viewQuery:function(t,o){if(1&t&&(e.Gf(q,7,e.Rgc),e.Gf(W,7,e.Rgc)),2&t){let s;e.iGM(s=e.CRH())&&(o.switchTpl=s.first),e.iGM(s=e.CRH())&&(o.copyTokenTpl=s.first)}},decls:11,vars:5,consts:[[1,""],[1,"list-header","pl-btnbase","pr-btnrbase","py-btnybase","flex","flex-nowrap","items-center","justify-between"],["eo-ng-button","","eoNgUserAccess","system/ext-app",1,"ant-btn-primary",3,"disabledEdit","click"],[1,"list-content"],[1,"",3,"nzTbody","nzThead","nzData","nzTrClick","nzMaxOperatorButton"],["copyTokenTpl",""],["switchTpl",""],["eo-copy","","copyType","click",1,"ant-btn","eo-ng-table-btn","eo-ng-table-delete-btn","ant-btn-text",2,"display","inline-flex","justify-content","center","align-items","center","vertical-align","middle",3,"copyText","copyCallback","click"],["eoNgFeedbackTooltip","","nzTooltipTitle","\u590d\u5236Token","nzTooltipPlacement","top","nzTooltipTrigger","hover",1,"inline-flex","items-center",3,"nzTooltipVisible"],[1,"iconpark-icon"],["href","#copy"],["eoNgUserAccess","system/ext-app",3,"ngModel","nzControl","nzDisabled","ngModelChange","click"]],template:function(t,o){1&t&&(e.TgZ(0,"div",0)(1,"div",1)(2,"div")(3,"button",2),e.NdJ("disabledEdit",function(r){return o.disabledEdit(r)})("click",function(){return o.addApp()}),e._uU(4," \u65b0\u5efa\u5e94\u7528 "),e.qZA()()(),e.TgZ(5,"div",3),e._UZ(6,"eo-ng-apinto-table",4),e.qZA()(),e.YNc(7,H,4,2,"ng-template",null,5,e.W1O),e.YNc(9,R,1,3,"ng-template",null,6,e.W1O)),2&t&&(e.xp6(6),e.Q6J("nzTbody",o.appsTableBody)("nzThead",o.appsTableHeadName)("nzData",o.appsList)("nzTrClick",o.appsTableClick)("nzMaxOperatorButton",4))},dependencies:[E.a,a.JJ,a.On,x.q,f.qK,y.g,Z.jw,c.Cl],encapsulation:2}),n})();var V=l(8513);let G=(()=>{class n{constructor(t,o,s,r){this.baseInfo=t,this.api=o,this.router=s,this.activateInfo=r,this.nowUrl=this.router.routerState.snapshot.url,this.appId="",this.baseInfo.allParamsInfo.extAppId?this.appId=this.baseInfo.allParamsInfo.extAppId:this.router.navigate(["/"])}ngOnDestroy(){}}return n.\u0275fac=function(t){return new(t||n)(e.Y36(V.D),e.Y36(h.s),e.Y36(p.F0),e.Y36(p.gz))},n.\u0275cmp=e.Xpm({type:n,selectors:[["eo-ng-external-app-message"]],decls:1,vars:2,consts:[[3,"editPage","appId"]],template:function(t,o){1&t&&e._UZ(0,"eo-ng-external-app-create",0),2&t&&e.Q6J("editPage",!0)("appId",o.appId)},dependencies:[_],encapsulation:2}),n})(),j=(()=>{class n{}return n.\u0275fac=function(t){return new(t||n)},n.\u0275cmp=e.Xpm({type:n,selectors:[["eo-ng-system"]],decls:1,vars:0,template:function(t,o){1&t&&e._UZ(0,"router-outlet")},dependencies:[p.lC],encapsulation:2}),n})();const K=[{title:"webhook\u540d\u79f0"},{title:"\u901a\u77e5url"},{title:"\u8bf7\u6c42\u65b9\u5f0f"},{title:"\u53c2\u6570\u7c7b\u578b"},{title:"\u66f4\u65b0\u8005"},{title:"\u66f4\u65b0\u65f6\u95f4"}],X=[{key:"title",copy:!0},{key:"url"},{key:"method"},{key:"contentType"},{key:"operator"},{key:"updateTime"}],$=[{key:"key",type:"input",placeholder:"\u8bf7\u8f93\u5165Key"},{key:"value",type:"input",placeholder:"\u8bf7\u8f93\u5165Value"}];var ee=l(8585);function te(n,i){1&n&&(e.ynx(0),e._uU(1,"\u4ec5\u652f\u6301\u4e2d\u82f1\u6587"),e.BQk())}function oe(n,i){1&n&&e.YNc(0,te,2,0,"ng-container",14),2&n&&e.Q6J("ngIf",i.$implicit.hasError("pattern"))}function ne(n,i){1&n&&(e.TgZ(0,"nz-form-item")(1,"nz-form-label",6)(2,"span",21),e._uU(3,"\u7528\u6237\u5206\u9694\u7b26\uff1a"),e.qZA()(),e.TgZ(4,"nz-form-control",6),e._UZ(5,"input",22),e.qZA()()),2&n&&(e.xp6(1),e.Q6J("nzSpan",7),e.xp6(1),e.Q6J("nzTooltipVisible",!1),e.xp6(2),e.Q6J("nzSpan",12))}let B=(()=>{class n{constructor(t,o,s){this.message=t,this.api=o,this.fb=s,this.webhookId="",this.autoTips=m.Sm,this.nzDisabled=!1,this.validateForm=new a.cw({}),this.methodsList=D,this.contentTypesList=L,this.noticeTypesList=U,this.responseHeaderList=[{key:"",value:""}],this.responseHeaderTableBody=[...$],this.validateForm=this.fb.group({title:["",[a.kI.required,ee.C.maxByteLength(32),a.kI.pattern("^[\u4e00-\u9fa5A-Za-z]+$")]],desc:[""],url:["",[a.kI.required]],method:["POST",[a.kI.required]],contentType:["JSON",[a.kI.required]],noticeType:["single",[a.kI.required]],userSeparator:[""],header:[""],template:["{}"]})}ngOnInit(){for(const t of this.responseHeaderTableBody)t.disabledFn=()=>this.nzDisabled;this.responseHeaderTableBody.push({type:"btn",showFn:t=>t===this.responseHeaderList[0],btns:[{title:"\u6dfb\u52a0",action:"add",disabledFn:()=>this.nzDisabled}]}),this.responseHeaderTableBody.push({type:"btn",showFn:t=>t!==this.responseHeaderList[0],btns:[{title:"\u6dfb\u52a0",action:"add",disabledFn:()=>this.nzDisabled},{title:"\u51cf\u5c11",action:"delete",disabledFn:()=>this.nzDisabled}]}),this.webhookId&&this.getWebhookMessage(this.webhookId)}disabledEdit(t){this.nzDisabled=t}getWebhookMessage(t){this.api.get("webhook",{uuid:t}).subscribe(o=>{0===o.code&&((0,g.m)(this.validateForm,o.data.webhook),this.validateForm.controls.template.setValue(o.data.webhook.template||"{}"),this.responseHeaderList=this.transferToList(o.data.webhook.header))})}transferToList(t){const o=[],s=Object.keys(t);if(s?.length>0){for(const r of s)o.push({key:r,value:t[r]});return o}return[{key:"",value:""}]}saveWebhook(){if(this.validateForm.valid){const t={...this.validateForm.value,header:this.transferToMap(this.responseHeaderList)};"single"===t.noticeType&&delete t.userSeparator,this.webhookId?this.api.put("webhook",{...t,uuid:this.webhookId}).subscribe(o=>{0===o.code&&(this.closeModal&&this.closeModal(),this.message.success(o.msg||"\u4fee\u6539Webhook\u6210\u529f\uff01"))}):this.api.post("webhook",t).subscribe(o=>{0===o.code&&(this.closeModal&&this.closeModal(),this.message.success(o.msg||"\u65b0\u5efaWebhook\u6210\u529f\uff01"))})}else Object.values(this.validateForm.controls).forEach(t=>{t.invalid&&(t.markAsDirty(),t.updateValueAndValidity({onlySelf:!0}))});return!1}transferToMap(t){const o={};for(const s of t)s.key&&s.value&&(o[s.key]=s.value);return o}}return n.\u0275fac=function(t){return new(t||n)(e.Y36(C.A),e.Y36(h.s),e.Y36(a.QS))},n.\u0275cmp=e.Xpm({type:n,selectors:[["eo-ng-system-webhook-config"]],inputs:{webhookId:"webhookId",closeModal:"closeModal"},decls:47,vars:31,consts:[["nz-form","","nzLayout","horizontal","autocomplete","off",3,"formGroup","nzAutoTips","nzNoColon"],["ngForm","ngForm"],["nzRequired","",3,"nzSpan"],["nzExtra","\u4ec5\u652f\u6301\u4e2d\u82f1\u6587\uff0c\u957f\u5ea632\u5b57\u7b26\u4ee5\u5185",3,"nzSpan","nzErrorTip"],["eo-ng-input","","formControlName","title","placeholder","\u8bf7\u8f93\u5165","eoNgUserAccess","system/email",1,"w-INPUT_NORMAL"],["nameErrorTpl",""],[3,"nzSpan"],["eo-ng-input","","formControlName","desc","placeholder","\u8bf7\u8f93\u5165","eoNgUserAccess","system/email",1,"w-INPUT_NORMAL"],["nzExtra","\u4ec5\u652f\u6301HTTP/HTTPS\u534f\u8baeAPI"],["eo-ng-input","","formControlName","url","placeholder","\u8bf7\u8f93\u5165\u901a\u77e5URL","eoNgUserAccess","system/email",1,"w-INPUT_NORMAL",3,"disabledEdit"],["formControlName","method","nzPlaceHolder","\u8bf7\u9009\u62e9",1,"w-INPUT_NORMAL",3,"nzOptions","nzDisabled"],["formControlName","contentType","nzPlaceHolder","\u8bf7\u9009\u62e9",1,"w-INPUT_NORMAL",3,"nzOptions","nzDisabled"],["eoNgFeedbackTooltip","","nzTooltipTitle","\u5355\u6b21\u53d1\u9001\u662f\u6307\u7f51\u5173\u6bcf\u89e6\u53d1\u4e00\u6b21\u544a\u8b66\u5c31\u4f1a\u8c03\u7528\u8be5\u63a5\u53e3\u4e00\u6b21\uff0c\u8be5\u63a5\u53e3\u5e94\u8be5\u652f\u6301\u7fa4\u53d1\u6d88\u606f\u7ed9\u7528\u6237\uff1b\u591a\u6b21\u53d1\u9001\u662f\u6307\u7f51\u5173\u6bcf\u89e6\u53d1\u4e00\u6b21\u544a\u8b66\u5c31\u4f1a\u6309\u7528\u6237\u4e2a\u6570\u8c03\u7528\u8be5\u63a5\u53e3\u3002","nzTooltipTrigger","hover",3,"nzTooltipVisible"],["formControlName","noticeType","nzPlaceHolder","\u8bf7\u9009\u62e9",1,"w-INPUT_NORMAL",3,"nzOptions","nzDisabled"],[4,"ngIf"],[1,"max-w-[380px]",3,"nzSpan"],[1,""],[1,"arrayItem","fuse-header",3,"nzTbody","nzData","nzNoScroll"],[1,"mb-0"],["nzExtra","\u63d0\u4f9b{title}\u3001{msg}\u3001{users}\u4e09\u4e2a\u53c2\u6570\u53d8\u91cf",3,"nzSpan"],["eo-ng-input","","formControlName","template","placeholder","\u8bf7\u8f93\u5165","eoNgUserAccess","system/email",1,"w-INPUT_NORMAL"],["eoNgFeedbackTooltip","","nzTooltipTitle","users\u53d8\u91cf\u503c\u4e2duser\u95f4\u7684\u5206\u9694\u7b26\uff0c\u6bd4\u5982\uff1auser1,user2\uff0c\u4f7f\u7528\u82f1\u6587\u2018,\u2019\u5206\u9694\u7b26\u3002","nzTooltipTrigger","hover",3,"nzTooltipVisible"],["eo-ng-input","","formControlName","userSeparator","placeholder","\u8bf7\u8f93\u5165","eoNgUserAccess","system/email",1,"w-INPUT_NORMAL"]],template:function(t,o){if(1&t&&(e.TgZ(0,"form",0,1)(2,"nz-form-item")(3,"nz-form-label",2),e._uU(4,"\u6a21\u677f\u540d\u79f0\uff1a"),e.qZA(),e.TgZ(5,"nz-form-control",3),e._UZ(6,"input",4),e.YNc(7,oe,1,1,"ng-template",null,5,e.W1O),e.qZA()(),e.TgZ(9,"nz-form-item")(10,"nz-form-label",6),e._uU(11,"\u63cf\u8ff0\uff1a"),e.qZA(),e.TgZ(12,"nz-form-control",6),e._UZ(13,"textarea",7),e.qZA()(),e.TgZ(14,"nz-form-item",8)(15,"nz-form-label",2),e._uU(16,"\u901a\u77e5URL\uff1a"),e.qZA(),e.TgZ(17,"nz-form-control",6)(18,"input",9),e.NdJ("disabledEdit",function(r){return o.disabledEdit(r)}),e.qZA()()(),e.TgZ(19,"nz-form-item")(20,"nz-form-label",2),e._uU(21,"\u8bf7\u6c42\u65b9\u5f0f\uff1a"),e.qZA(),e.TgZ(22,"nz-form-control",6),e._UZ(23,"eo-ng-select",10),e.qZA()(),e.TgZ(24,"nz-form-item")(25,"nz-form-label",2),e._uU(26,"\u53c2\u6570\u7c7b\u578b\uff1a"),e.qZA(),e.TgZ(27,"nz-form-control",6),e._UZ(28,"eo-ng-select",11),e.qZA()(),e.TgZ(29,"nz-form-item")(30,"nz-form-label",2)(31,"span",12),e._uU(32,"\u6d88\u606f\u7c7b\u578b\uff1a "),e.qZA()(),e.TgZ(33,"nz-form-control",6),e._UZ(34,"eo-ng-select",13),e.qZA()(),e.YNc(35,ne,6,3,"nz-form-item",14),e.TgZ(36,"nz-form-item")(37,"nz-form-label",6),e._uU(38,"Header\u53c2\u6570\uff1a"),e.qZA(),e.TgZ(39,"nz-form-control",15)(40,"div",16),e._UZ(41,"eo-ng-apinto-table",17),e.qZA()()(),e.TgZ(42,"nz-form-item",18)(43,"nz-form-label",6),e._uU(44,"\u53c2\u6570\u6a21\u677f\uff1a"),e.qZA(),e.TgZ(45,"nz-form-control",19),e._UZ(46,"textarea",20),e.qZA()()()),2&t){const s=e.MAs(8);e.Q6J("formGroup",o.validateForm)("nzAutoTips",o.autoTips)("nzNoColon",!0),e.xp6(3),e.Q6J("nzSpan",7),e.xp6(2),e.Q6J("nzSpan",12)("nzErrorTip",s),e.xp6(5),e.Q6J("nzSpan",7),e.xp6(2),e.Q6J("nzSpan",12),e.xp6(3),e.Q6J("nzSpan",7),e.xp6(2),e.Q6J("nzSpan",12),e.xp6(3),e.Q6J("nzSpan",7),e.xp6(2),e.Q6J("nzSpan",12),e.xp6(1),e.Q6J("nzOptions",o.methodsList)("nzDisabled",o.nzDisabled),e.xp6(2),e.Q6J("nzSpan",7),e.xp6(2),e.Q6J("nzSpan",12),e.xp6(1),e.Q6J("nzOptions",o.contentTypesList)("nzDisabled",o.nzDisabled),e.xp6(2),e.Q6J("nzSpan",7),e.xp6(1),e.Q6J("nzTooltipVisible",!1),e.xp6(2),e.Q6J("nzSpan",12),e.xp6(1),e.Q6J("nzOptions",o.noticeTypesList)("nzDisabled",o.nzDisabled),e.xp6(1),e.Q6J("ngIf","single"===(null==o.validateForm.controls.noticeType?null:o.validateForm.controls.noticeType.value)),e.xp6(2),e.Q6J("nzSpan",7),e.xp6(2),e.Q6J("nzSpan",12),e.xp6(2),e.Q6J("nzTbody",o.responseHeaderTableBody)("nzData",o.responseHeaderList)("nzNoScroll",!0),e.xp6(2),e.Q6J("nzSpan",7),e.xp6(2),e.Q6J("nzSpan",12)}},dependencies:[A.O5,E.a,a._Y,a.Fj,a.JJ,a.JL,a.sg,a.u,F.oO,b.t3,b.SK,u.Lr,u.Nx,u.iK,u.Fd,y.g,v.Hf,c.Cl],encapsulation:2}),n})();function se(n,i){if(1&n&&e._UZ(0,"eo-ng-system-webhook-config",8,9),2&n){const t=e.oxw();e.Q6J("webhookId",t.webhookId)}}const ae=[{path:"",component:j,data:{id:"6"},children:[{path:"ext-app",component:P,data:{id:"602"},children:[{path:"",component:Y},{path:"create",component:_},{path:"message/:extAppId",component:G}]},{path:"email",component:Q,data:{id:"603"},children:[{path:"",component:O}]},{path:"webhook",component:(()=>{class n{}return n.\u0275fac=function(t){return new(t||n)},n.\u0275cmp=e.Xpm({type:n,selectors:[["eo-ng-system-webhook"]],decls:1,vars:0,template:function(t,o){1&t&&e._UZ(0,"router-outlet")},dependencies:[p.lC],encapsulation:2}),n})(),data:{id:"604"},children:[{path:"",component:(()=>{class n{constructor(t,o,s,r){this.message=t,this.modalService=o,this.api=s,this.appConfigService=r,this.webhooksList=[],this.nzDisabled=!1,this.webhooksTableHead=[...K],this.webhooksTableBody=[...X],this.webhookId="",this.webhookTableClick=d=>{this.openWebhookModal(d.data.uuid)},this.closeModal=()=>{this.getWebhooksList(),this.modalRef?.close()},this.appConfigService.reqFlashBreadcrumb([{title:"Webhook\u7ba1\u7406",routerLink:"system/webhook"}])}ngOnInit(){this.webhooksTableHead.push({title:"\u64cd\u4f5c",right:!0}),this.webhooksTableBody.push({type:"btn",right:!0,btns:[{title:"\u67e5\u770b",click:t=>{this.openWebhookModal(t.data.uuid)}},{title:"\u5220\u9664",disabledFn:(t,o)=>!o.data.isDelete||this.nzDisabled,click:t=>{this.delete(t.data)}}]}),this.getWebhooksList()}disabledEdit(t){this.nzDisabled=t}getWebhooksList(){this.api.get("webhooks").subscribe(t=>{0===t.code&&(this.webhooksList=t.data.webhooks)})}openWebhookModal(t=""){this.webhookId=t||"",this.modalRef=this.modalService.create({nzTitle:this.webhookId?"\u65b0\u5efaWebhook":"\u7f16\u8f91Webhook",nzContent:B,nzComponentParams:{webhookId:this.webhookId,closeModal:this.closeModal},nzClosable:!0,nzWidth:S.iE,nzOkText:this.webhookId?"\u63d0\u4ea4":"\u4fdd\u5b58",nzOnOk:o=>(o.saveWebhook(),!1)})}delete(t,o){o?.stopPropagation(),this.modalService.create({nzTitle:"\u5220\u9664",nzContent:`${t.title}\u4e00\u65e6\u5220\u9664\uff0c\u6570\u636e\u5c06\u4f1a\u4e22\u5931\u3002`,nzClosable:!0,nzWidth:S.G_,nzClassName:"delete-modal",nzOkDanger:!0,nzOnOk:()=>{this.deleteDiscovery(t.uuid)}})}deleteDiscovery(t){this.api.delete("webhook",{uuid:t}).subscribe(o=>{0===o.code&&(this.getWebhooksList(),this.message.success(o.msg||"\u5220\u9664\u6210\u529f"))})}}return n.\u0275fac=function(t){return new(t||n)(e.Y36(C.A),e.Y36(c._B),e.Y36(h.s),e.Y36(k.a))},n.\u0275cmp=e.Xpm({type:n,selectors:[["eo-ng-system-webhook-list"]],decls:11,vars:6,consts:[[1,"pl-btnbase","pr-btnrbase","py-btnybase","flex","flex-nowrap","items-center","justify-between"],["eo-ng-button","","nzType","primary","eoNgUserAccess","system/email",1,"ant-btn-primary",3,"click"],[1,"iconpark-icon"],["href","#add"],[1,"ml-label"],[1,"list-content","pb-[4px]"],[3,"nzTbody","nzThead","nzData","nzTrClick","nzMaxOperatorButton","nzNoScroll"],["webhookDetailTpl",""],[3,"webhookId"],["webhookFormRef",""]],template:function(t,o){1&t&&(e.TgZ(0,"div")(1,"div",0)(2,"button",1),e.NdJ("click",function(){return o.openWebhookModal()}),e.O4$(),e.TgZ(3,"svg",2),e._UZ(4,"use",3),e.qZA(),e.kcU(),e.TgZ(5,"span",4),e._uU(6," \u65b0\u5efaWebhook"),e.qZA()()(),e.TgZ(7,"div",5),e._UZ(8,"eo-ng-apinto-table",6),e.qZA()(),e.YNc(9,se,2,1,"ng-template",null,7,e.W1O)),2&t&&(e.xp6(8),e.Q6J("nzTbody",o.webhooksTableBody)("nzThead",o.webhooksTableHead)("nzData",o.webhooksList)("nzTrClick",o.webhookTableClick)("nzMaxOperatorButton",2)("nzNoScroll",!0))},dependencies:[E.a,f.qK,y.g,B],encapsulation:2}),n})()}]}]}];let ie=(()=>{class n{}return n.\u0275fac=function(t){return new(t||n)},n.\u0275mod=e.oAB({type:n}),n.\u0275inj=e.cJS({imports:[p.Bz.forChild(ae),p.Bz]}),n})();var le=l(3111),re=l(2732),pe=l(7901);let ue=(()=>{class n{}return n.\u0275fac=function(t){return new(t||n)},n.\u0275mod=e.oAB({type:n}),n.\u0275inj=e.cJS({imports:[A.ez,ie,le.a,re.D,a.u5,a.UX,x.r,F.FC,f.r1,u.U5,pe.Z,Z.WF,v.HO,c.mx,c.Jy]}),n})()}}]);