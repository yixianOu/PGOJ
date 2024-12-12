import{m as R,d as be,c as C,a as M,E as Ee,M as Y,w as _,o as Be,b as we,ae as Me,f as k,g as c,h as v,j as m,P as F,k as b,A as p,x as r,D as x,C as z,q as K,B as P,y as S,J as L,I as D,r as ke,H as N,U as j,T as Se,W as ie,a7 as ue,Z as Te,a8 as ce,af as ee,$ as ne}from"./index-BxSgbHR9.js";import{E as Ie,a as Ae}from"./el-input-B9UNxcP1.js";import{E as Le}from"./index-BSl0X_TB.js";import{E as Re,u as $e,a as Ve,b as Oe}from"./index-BvgBipFQ.js";import{d as ze,E as Pe,aS as De,aB as oe,ag as se,a as ae,p as He,e as Ue,H as te}from"./constants-BbhSUWfP.js";import{_ as Fe}from"./base-D0W_z4C0.js";import{E as Ke}from"./aria-BUADUvnR.js";import{o as le}from"./aria-nkjrUMQ-.js";import{E as Ne}from"./focus-trap-t7Ag-kVD.js";const je=e=>["",...ze].includes(e),q="_trap-focus-children",E=[],re=e=>{if(E.length===0)return;const o=E[E.length-1][q];if(o.length>0&&e.code===Ke.tab){if(o.length===1){e.preventDefault(),document.activeElement!==o[0]&&o[0].focus();return}const t=e.shiftKey,i=e.target===o[0],l=e.target===o[o.length-1];i&&t&&(e.preventDefault(),o[o.length-1].focus()),l&&!t&&(e.preventDefault(),o[0].focus())}},qe={beforeMount(e){e[q]=le(e),E.push(e),E.length<=1&&document.addEventListener("keydown",re)},updated(e){R(()=>{e[q]=le(e)})},unmounted(){E.shift(),E.length===0&&document.removeEventListener("keydown",re)}},Ge=be({name:"ElMessageBox",directives:{TrapFocus:qe},components:{ElButton:Ie,ElFocusTrap:Ne,ElInput:Le,ElOverlay:Re,ElIcon:Pe,...De},inheritAttrs:!1,props:{buttonSize:{type:String,validator:je},modal:{type:Boolean,default:!0},lockScroll:{type:Boolean,default:!0},showClose:{type:Boolean,default:!0},closeOnClickModal:{type:Boolean,default:!0},closeOnPressEscape:{type:Boolean,default:!0},closeOnHashChange:{type:Boolean,default:!0},center:Boolean,draggable:Boolean,overflow:Boolean,roundButton:{default:!1,type:Boolean},container:{type:String,default:"body"},boxType:{type:String,default:""}},emits:["vanish","action"],setup(e,{emit:o}){const{locale:t,zIndex:i,ns:l,size:s}=Ae("message-box",C(()=>e.buttonSize)),{t:u}=t,{nextZIndex:f}=i,y=M(!1),n=Ee({autofocus:!0,beforeClose:null,callback:null,cancelButtonText:"",cancelButtonClass:"",confirmButtonText:"",confirmButtonClass:"",customClass:"",customStyle:{},dangerouslyUseHTMLString:!1,distinguishCancelAndClose:!1,icon:"",inputPattern:null,inputPlaceholder:"",inputType:"text",inputValue:null,inputValidator:null,inputErrorMessage:"",message:null,modalFade:!0,modalClass:"",showCancelButton:!1,showConfirmButton:!0,type:"",title:void 0,showInput:!1,action:"",confirmButtonLoading:!1,cancelButtonLoading:!1,confirmButtonLoadingIcon:Y(oe),cancelButtonLoadingIcon:Y(oe),confirmButtonDisabled:!1,editorErrorMessage:"",validateError:!1,zIndex:f()}),H=C(()=>{const a=n.type;return{[l.bm("icon",a)]:a&&se[a]}}),U=ae(),d=ae(),de=C(()=>n.icon||se[n.type]||""),fe=C(()=>!!n.message),B=M(),G=M(),I=M(),V=M(),W=M(),me=C(()=>n.confirmButtonClass);_(()=>n.inputValue,async a=>{await R(),e.boxType==="prompt"&&a!==null&&Z()},{immediate:!0}),_(()=>y.value,a=>{var g,w;a&&(e.boxType!=="prompt"&&(n.autofocus?I.value=(w=(g=W.value)==null?void 0:g.$el)!=null?w:B.value:I.value=B.value),n.zIndex=f()),e.boxType==="prompt"&&(a?R().then(()=>{var Q;V.value&&V.value.$el&&(n.autofocus?I.value=(Q=ye())!=null?Q:B.value:I.value=B.value)}):(n.editorErrorMessage="",n.validateError=!1))});const pe=C(()=>e.draggable),ge=C(()=>e.overflow);$e(B,G,pe,ge),Be(async()=>{await R(),e.closeOnHashChange&&window.addEventListener("hashchange",A)}),we(()=>{e.closeOnHashChange&&window.removeEventListener("hashchange",A)});function A(){y.value&&(y.value=!1,R(()=>{n.action&&o("action",n.action)}))}const X=()=>{e.closeOnClickModal&&O(n.distinguishCancelAndClose?"close":"cancel")},ve=Oe(X),he=a=>{if(n.inputType!=="textarea")return a.preventDefault(),O("confirm")},O=a=>{var g;e.boxType==="prompt"&&a==="confirm"&&!Z()||(n.action=a,n.beforeClose?(g=n.beforeClose)==null||g.call(n,a,n,A):A())},Z=()=>{if(e.boxType==="prompt"){const a=n.inputPattern;if(a&&!a.test(n.inputValue||""))return n.editorErrorMessage=n.inputErrorMessage||u("el.messagebox.error"),n.validateError=!0,!1;const g=n.inputValidator;if(typeof g=="function"){const w=g(n.inputValue);if(w===!1)return n.editorErrorMessage=n.inputErrorMessage||u("el.messagebox.error"),n.validateError=!0,!1;if(typeof w=="string")return n.editorErrorMessage=w,n.validateError=!0,!1}}return n.editorErrorMessage="",n.validateError=!1,!0},ye=()=>{const a=V.value.$refs;return a.input||a.textarea},J=()=>{O("close")},Ce=()=>{e.closeOnPressEscape&&J()};return e.lockScroll&&Ve(y),{...Me(n),ns:l,overlayEvent:ve,visible:y,hasMessage:fe,typeClass:H,contentId:U,inputId:d,btnSize:s,iconComponent:de,confirmButtonClasses:me,rootRef:B,focusStartRef:I,headerRef:G,inputRef:V,confirmRef:W,doClose:A,handleClose:J,onCloseRequested:Ce,handleWrapperClick:X,handleInputEnter:he,handleAction:O,t:u}}});function We(e,o,t,i,l,s){const u=k("el-icon"),f=k("close"),y=k("el-input"),n=k("el-button"),H=k("el-focus-trap"),U=k("el-overlay");return c(),v(Se,{name:"fade-in-linear",onAfterLeave:d=>e.$emit("vanish"),persisted:""},{default:m(()=>[F(b(U,{"z-index":e.zIndex,"overlay-class":[e.ns.is("message-box"),e.modalClass],mask:e.modal},{default:m(()=>[p("div",{role:"dialog","aria-label":e.title,"aria-modal":"true","aria-describedby":e.showInput?void 0:e.contentId,class:r(`${e.ns.namespace.value}-overlay-message-box`),onClick:e.overlayEvent.onClick,onMousedown:e.overlayEvent.onMousedown,onMouseup:e.overlayEvent.onMouseup},[b(H,{loop:"",trapped:e.visible,"focus-trap-el":e.rootRef,"focus-start-el":e.focusStartRef,onReleaseRequested:e.onCloseRequested},{default:m(()=>[p("div",{ref:"rootRef",class:r([e.ns.b(),e.customClass,e.ns.is("draggable",e.draggable),{[e.ns.m("center")]:e.center}]),style:x(e.customStyle),tabindex:"-1",onClick:z(()=>{},["stop"])},[e.title!==null&&e.title!==void 0?(c(),K("div",{key:0,ref:"headerRef",class:r([e.ns.e("header"),{"show-close":e.showClose}])},[p("div",{class:r(e.ns.e("title"))},[e.iconComponent&&e.center?(c(),v(u,{key:0,class:r([e.ns.e("status"),e.typeClass])},{default:m(()=>[(c(),v(P(e.iconComponent)))]),_:1},8,["class"])):S("v-if",!0),p("span",null,L(e.title),1)],2),e.showClose?(c(),K("button",{key:0,type:"button",class:r(e.ns.e("headerbtn")),"aria-label":e.t("el.messagebox.close"),onClick:d=>e.handleAction(e.distinguishCancelAndClose?"close":"cancel"),onKeydown:D(z(d=>e.handleAction(e.distinguishCancelAndClose?"close":"cancel"),["prevent"]),["enter"])},[b(u,{class:r(e.ns.e("close"))},{default:m(()=>[b(f)]),_:1},8,["class"])],42,["aria-label","onClick","onKeydown"])):S("v-if",!0)],2)):S("v-if",!0),p("div",{id:e.contentId,class:r(e.ns.e("content"))},[p("div",{class:r(e.ns.e("container"))},[e.iconComponent&&!e.center&&e.hasMessage?(c(),v(u,{key:0,class:r([e.ns.e("status"),e.typeClass])},{default:m(()=>[(c(),v(P(e.iconComponent)))]),_:1},8,["class"])):S("v-if",!0),e.hasMessage?(c(),K("div",{key:1,class:r(e.ns.e("message"))},[ke(e.$slots,"default",{},()=>[e.dangerouslyUseHTMLString?(c(),v(P(e.showInput?"label":"p"),{key:1,for:e.showInput?e.inputId:void 0,innerHTML:e.message},null,8,["for","innerHTML"])):(c(),v(P(e.showInput?"label":"p"),{key:0,for:e.showInput?e.inputId:void 0},{default:m(()=>[N(L(e.dangerouslyUseHTMLString?"":e.message),1)]),_:1},8,["for"]))])],2)):S("v-if",!0)],2),F(p("div",{class:r(e.ns.e("input"))},[b(y,{id:e.inputId,ref:"inputRef",modelValue:e.inputValue,"onUpdate:modelValue":d=>e.inputValue=d,type:e.inputType,placeholder:e.inputPlaceholder,"aria-invalid":e.validateError,class:r({invalid:e.validateError}),onKeydown:D(e.handleInputEnter,["enter"])},null,8,["id","modelValue","onUpdate:modelValue","type","placeholder","aria-invalid","class","onKeydown"]),p("div",{class:r(e.ns.e("errormsg")),style:x({visibility:e.editorErrorMessage?"visible":"hidden"})},L(e.editorErrorMessage),7)],2),[[j,e.showInput]])],10,["id"]),p("div",{class:r(e.ns.e("btns"))},[e.showCancelButton?(c(),v(n,{key:0,loading:e.cancelButtonLoading,"loading-icon":e.cancelButtonLoadingIcon,class:r([e.cancelButtonClass]),round:e.roundButton,size:e.btnSize,onClick:d=>e.handleAction("cancel"),onKeydown:D(z(d=>e.handleAction("cancel"),["prevent"]),["enter"])},{default:m(()=>[N(L(e.cancelButtonText||e.t("el.messagebox.cancel")),1)]),_:1},8,["loading","loading-icon","class","round","size","onClick","onKeydown"])):S("v-if",!0),F(b(n,{ref:"confirmRef",type:"primary",loading:e.confirmButtonLoading,"loading-icon":e.confirmButtonLoadingIcon,class:r([e.confirmButtonClasses]),round:e.roundButton,disabled:e.confirmButtonDisabled,size:e.btnSize,onClick:d=>e.handleAction("confirm"),onKeydown:D(z(d=>e.handleAction("confirm"),["prevent"]),["enter"])},{default:m(()=>[N(L(e.confirmButtonText||e.t("el.messagebox.confirm")),1)]),_:1},8,["loading","loading-icon","class","round","disabled","size","onClick","onKeydown"]),[[j,e.showConfirmButton]])],2)],14,["onClick"])]),_:3},8,["trapped","focus-trap-el","focus-start-el","onReleaseRequested"])],42,["aria-label","aria-describedby","onClick","onMousedown","onMouseup"])]),_:3},8,["z-index","overlay-class","mask"]),[[j,e.visible]])]),_:3},8,["onAfterLeave"])}var Xe=Fe(Ge,[["render",We],["__file","index.vue"]]);const $=new Map,Ze=e=>{let o=document.body;return e.appendTo&&(ie(e.appendTo)&&(o=document.querySelector(e.appendTo)),te(e.appendTo)&&(o=e.appendTo),te(o)||(o=document.body)),o},Je=(e,o,t=null)=>{const i=b(Xe,e,ne(e.message)||ue(e.message)?{default:ne(e.message)?e.message:()=>e.message}:null);return i.appContext=t,ce(i,o),Ze(e).appendChild(o.firstElementChild),i.component},Qe=()=>document.createElement("div"),Ye=(e,o)=>{const t=Qe();e.onVanish=()=>{ce(null,t),$.delete(l)},e.onAction=s=>{const u=$.get(l);let f;e.showInput?f={value:l.inputValue,action:s}:f=s,e.callback?e.callback(f,i.proxy):s==="cancel"||s==="close"?e.distinguishCancelAndClose&&s!=="cancel"?u.reject("close"):u.reject("cancel"):u.resolve(f)};const i=Je(e,t,o),l=i.proxy;for(const s in e)ee(e,s)&&!ee(l.$props,s)&&(l[s]=e[s]);return l.visible=!0,l};function T(e,o=null){if(!He)return Promise.reject();let t;return ie(e)||ue(e)?e={message:e}:t=e.callback,new Promise((i,l)=>{const s=Ye(e,o??T._context);$.set(s,{options:e,callback:t,resolve:i,reject:l})})}const _e=["alert","confirm","prompt"],xe={alert:{closeOnPressEscape:!1,closeOnClickModal:!1},confirm:{showCancelButton:!0},prompt:{showCancelButton:!0,showInput:!0}};_e.forEach(e=>{T[e]=en(e)});function en(e){return(o,t,i,l)=>{let s="";return Te(t)?(i=t,s=""):Ue(t)?s="":s=t,T(Object.assign({title:s,message:o,type:"",...xe[e]},i,{boxType:e}),l)}}T.close=()=>{$.forEach((e,o)=>{o.doClose()}),$.clear()};T._context=null;const h=T;h.install=e=>{h._context=e._context,e.config.globalProperties.$msgbox=h,e.config.globalProperties.$messageBox=h,e.config.globalProperties.$alert=h.alert,e.config.globalProperties.$confirm=h.confirm,e.config.globalProperties.$prompt=h.prompt};const dn=h;export{dn as E};